package main

import (
	"github.com/mostlyworks/FetchReceiptService/api"
)

func main() {
	var server = api.CreateServer()

	api.StartServer(server)
}
