package services

import (
	"regexp"
	"strings"

	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/shopspring/decimal"
)

var pointConfig models.PointConfig

func InitPointsService(appPointConfig models.PointConfig) {
	pointConfig = appPointConfig
}

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
func receiptTotalPoints(total decimal.Decimal) int {
	var zero = decimal.NewFromFloat(0)

	var points = 0
	// decimal.NewFromFloat(0) should be a constant but can't be.

	if total.Mod(decimal.NewFromFloat(pointConfig.TotalRoundMod)).Equal(zero) {
		points += pointConfig.TotalRoundedPoints
	}
	if total.Mod(decimal.NewFromFloat(pointConfig.TotalMutiple)).Equal(zero) {
		points += pointConfig.TotalMutiplePoints
	}

	if points == 0 {
		return pointConfig.DefaultPointReturn
	}

	return points
}

// One point for every alphanumeric character in the retailer name.
// Calculate configured points for cleaned retailer name, given retailer name string.
func retailerPoints(retailer string) int {
	// Cleaned retailer name of spaces and non alpha numeric characters,
	// This should probably include accented latin characters, I don't want want to mess around with regex more.
	// Does not support localization outside of EU/US. Probably to break on Cyrillic, kanji, hanja, ect.
	retailRegex := regexp.MustCompile(`[^a-zA-Z0-9\s\:]*`)
	// Regex isn't capturing spaces correctly, do it the old fashioned way.
	var cleanedRetailer = retailRegex.ReplaceAllString(strings.ReplaceAll(retailer, " ", ""), "")
	return len(cleanedRetailer) * pointConfig.RetailerNamePointMutiplier
}

// 5 points for every two items on the receipt.
// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// Calculate points for items from supplied array.
func itemPoints(stringItems []models.Item) int {
	var points = 0

	points += (len(stringItems) / pointConfig.ItemCountDivsor * pointConfig.ItemCountPoints)

	for index := range stringItems {
		var value = stringItems[index]

		if len(strings.Trim(value.ShortDescription, " "))%pointConfig.ItemDescriptionMutiple == 0 {

			points += int(value.Price.Mul(decimal.NewFromFloat(pointConfig.ItemDescriptionPriceMutiplier)).RoundUp(pointConfig.PriceMutiplierRoundingPoints).BigInt().Int64())
		}

	}

	if points == 0 {
		return pointConfig.DefaultPointReturn
	}

	return points
}

// 6 points if the day in the purchase date is odd.
// Calculate points for configured check of day from given date string
func datePoints(time models.Date) int {

	if time.Time.Day()%pointConfig.PurchaseDateCheckMod != 0 {
		return pointConfig.PurchaseDatePoints
	}

	return pointConfig.DefaultPointReturn

}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
// Calcuate points for configured time window from given time string
// This should probably consider Minutes as well instead of just being on the hours.
func timePoints(time models.Time) int {

	if time.Time.Hour() >= pointConfig.PurchaseTimeLowerBound && time.Hour() <= pointConfig.PurchaseTimeUpperBound {
		return pointConfig.PurchaseTimePoints
	}

	return pointConfig.DefaultPointReturn
}
