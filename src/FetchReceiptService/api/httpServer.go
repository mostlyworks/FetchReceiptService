package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/mostlyworks/FetchReceiptService/repository"
)

type App struct {
	httpServer       *http.ServeMux
	database         repository.Database
	httpServerConfig models.HttpServerConfig
	pointsConfig     models.PointConfig
	databaseConfig   models.DatabaseConfig
}

func CreateServer() App {

	app := App{
		database:         repository.Init(),
		httpServer:       http.NewServeMux(),
		pointsConfig:     loadPointConfig(),
		httpServerConfig: loadHttpConfig(),
	}

	ReceiptRouter(app)

	return app
}

func StartServer(app App) {

	fmt.Println("Server up on " + strconv.Itoa(app.httpServerConfig.Port))
	if app.httpServerConfig.UseHttps {
		log.Print("Serving HTTPS")
		// Probably better validation on the file locations would be good.
		if app.httpServerConfig.Cert == "" || app.httpServerConfig.Key == "" {
			log.Panic("HTTPS configuration issue, running in HTTP")
			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(app.httpServerConfig.Port), app.httpServer))
		} else {
			log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(app.httpServerConfig.Port), app.httpServerConfig.Cert, app.httpServerConfig.Key, app.httpServer))
		}
	} else {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(app.httpServerConfig.Port), app.httpServer))
	}
}

func loadPointConfig() models.PointConfig {
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

func loadHttpConfig() models.HttpServerConfig {
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
