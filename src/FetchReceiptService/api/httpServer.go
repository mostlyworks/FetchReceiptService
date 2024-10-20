package api

import (
	"fmt"
	"github.com/mostlyworks/FetchReceiptService/config"
	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/mostlyworks/FetchReceiptService/repository"
	"log"
	"net/http"
	"strconv"
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
		pointsConfig:     config.LoadPointConfig(),
		httpServerConfig: config.LoadHttpConfig(),
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
