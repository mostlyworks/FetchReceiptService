package services

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/shopspring/decimal"
)

const totalRoundPoints = 50
const totalMutiplePoints = 25
const totalMutiple = 0.25
const totalRoundMod = 1.00
const itemCountPoints = 5
const itemCountDivsor = 2
const itemDescriptionMutiple = 3
const itemDescriptionPriceMutiplier = 0.2
const priceMutiplierRoundingPoints = 0
const purchaseDatePoints = 6
const purchaseDateCheckMod = 2
const purchaseTimeLowerBound = 14
const purchaseTimeUpperBound = 16
const purchaseTimePoints = 10
const dateExpectedFormat = "2006-01-02"
const timeExpectedFormat = "15:04"
const defaultPointReturn = 0

//TODO: Make points configurable on startup via config json.

func GetPoints(receipt models.Receipt) int {
	var pointHold = 0

	pointHold += receiptTotalPoints(receipt.Total)
	pointHold += retailerPoints(receipt.Retailer)
	pointHold += itemPoints(receipt.Items)
	pointHold += datePoints(receipt.PurchaseDate)
	pointHold += timePoints(receipt.PurchaseTime)

	return pointHold
}

// 50 points if the total is a round dollar amount with no cents.
// 25 points if the total is a multiple of 0.25.
// Calculate points on given string total of receipt
func receiptTotalPoints(stringTotal string) int {
	total, err := decimal.NewFromString(stringTotal)
	if err != nil {
		// This should be covered by validation?
		log.Fatal(err)
	}

	var points = 0
	// decimal.NewFromFloat(0) should be a constant but can't be.

	if total.Mod(decimal.NewFromFloat(totalRoundMod)).Equal(decimal.NewFromFloat(0)) {
		points += totalRoundPoints
	}
	if total.Mod(decimal.NewFromFloat(totalMutiple)).Equal(decimal.NewFromFloat(0)) {
		points += totalMutiplePoints
	}

	if points == 0 {
		return defaultPointReturn
	}

	return points
}

// One point for every alphanumeric character in the retailer name.
// Calculate configured points for cleaned retailer name, given retailer name string.
func retailerPoints(retailer string) int {
	// Cleaned retailer name of spaces and non alpha numeric characters,
	// This should probably include accented latin characters, I don't want want to mess around with regex more.
	// Does not support localization outside of EU/US
	retailRegex := regexp.MustCompile(`[^a-zA-Z0-9\s\:]*`)
	// Regex isn't capturing spaces correctly, do it the old fashioned way.
	var cleanedRetailer = retailRegex.ReplaceAllString(strings.ReplaceAll(retailer, " ", ""), "")
	return len(cleanedRetailer)
}

// 5 points for every two items on the receipt.
// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// Calculate points for items from supplied array.
func itemPoints(stringItems []models.Item) int {
	var points = 0

	points += (len(stringItems) / itemCountDivsor * itemCountPoints)

	for index := range stringItems {
		var value = stringItems[index]

		if len(strings.Trim(value.ShortDescription, " "))%itemDescriptionMutiple == 0 {
			price, err := decimal.NewFromString(value.Price)
			if err != nil {
				// This should be covered by validation?
				log.Fatal(err)
			}

			points += int(price.Mul(decimal.NewFromFloat(itemDescriptionPriceMutiplier)).RoundUp(priceMutiplierRoundingPoints).BigInt().Int64())
		}

	}

	if points == 0 {
		return defaultPointReturn
	}

	return points
}

// 6 points if the day in the purchase date is odd.
// Calculate points for configured check of day from given date string
func datePoints(stringDate string) int {
	time, err := time.Parse(dateExpectedFormat, stringDate)

	if err != nil {
		// This should be covered by validation?
		log.Fatal(err)
	}

	if time.Day()%purchaseDateCheckMod != 0 {
		return purchaseDatePoints
	}

	return defaultPointReturn

}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
// Calcuate points for configured time window from given time string
// This should probably consider Minutes as well instead of just being on the hours.
func timePoints(stringTime string) int {
	time, err := time.Parse(timeExpectedFormat, stringTime)

	if err != nil {
		// This should be covered by validation?
		log.Fatal(err)
	}

	if time.Hour() >= purchaseTimeLowerBound && time.Hour() <= purchaseTimeUpperBound {
		return purchaseTimePoints
	}

	return defaultPointReturn
}
