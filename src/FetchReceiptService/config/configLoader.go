package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mostlyworks/FetchReceiptService/models"
)

func LoadPointConfig() models.PointConfig {
	var pointConfig models.PointConfig
	pointConfigFile, err := os.Open("./config/pointConfig.json")
	defer pointConfigFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		log.Print("Loading with Points Default config")
		// Load defaults instead
		pointConfig = loadDefaultPointConfig(pointConfig)
	} else {
		jsonParser := json.NewDecoder(pointConfigFile)
		jsonParser.Decode(&pointConfig)
		log.Print("Loaded Points Config")
	}

	return pointConfig
}

func LoadHttpConfig() models.HttpServerConfig {
	var httpServerConfig models.HttpServerConfig
	httpServerConfigFile, err := os.Open("./config/httpServerConfig.json")
	defer httpServerConfigFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		log.Print("Loading with HTTP Default config")
		// Load defaults instead
		httpServerConfig = loadDefaultHttpConfig(httpServerConfig)
	} else {
		jsonParser := json.NewDecoder(httpServerConfigFile)
		jsonParser.Decode(&httpServerConfig)
		log.Print("Loaded http Config")
	}

	return httpServerConfig
}

func loadDefaultPointConfig(pointConfig models.PointConfig) models.PointConfig {
	pointConfig.TotalRoundedPoints = 50
	pointConfig.TotalMutiplePoints = 25
	pointConfig.TotalMutiple = 0.25
	pointConfig.TotalRoundMod = 1.00
	pointConfig.ItemCountPoints = 5
	pointConfig.ItemCountDivsor = 2
	pointConfig.ItemDescriptionMutiple = 3
	pointConfig.ItemDescriptionPriceMutiplier = 0.2
	pointConfig.PriceMutiplierRoundingPoints = 0
	pointConfig.PurchaseDatePoints = 6
	pointConfig.PurchaseDateCheckMod = 2
	pointConfig.PurchaseTimeLowerBound = 14
	pointConfig.PurchaseTimeUpperBound = 16
	pointConfig.PurchaseTimePoints = 10
	pointConfig.DateExpectedFormat = "2006-01-02"
	pointConfig.TimeExpectedFormat = "15:04"
	pointConfig.RetailerNamePointMutiplier = 1
	pointConfig.DefaultPointReturn = 0

	return pointConfig
}

func loadDefaultHttpConfig(httpServerConfig models.HttpServerConfig) models.HttpServerConfig {
	httpServerConfig.UseHttps = false
	httpServerConfig.Port = 8080

	return httpServerConfig
}
