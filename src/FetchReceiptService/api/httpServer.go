package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mostlyworks/FetchReceiptService/repository"
)

type App struct {
	httpServer *http.ServeMux
	database   repository.Database
}

func CreateServer() App {

	app := App{
		database:   repository.Init(),
		httpServer: http.NewServeMux(),
	}

	ReceiptRouter(app)

	return app
}

func StartServer(app App) {

	fmt.Println("Server up on 8080")
	log.Fatal(http.ListenAndServe(":8080", app.httpServer))
}
