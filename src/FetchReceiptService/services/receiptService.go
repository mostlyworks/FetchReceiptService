package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/mostlyworks/FetchReceiptService/repository"
)

var database repository.Database

func InitReceiptService(serverDB repository.Database) {
	database = serverDB
}

func WriteReciept(receipt models.Receipt) uuid.UUID {

	// If this was a real database, CLEAN THE DATA BEFORE INSERTION.
	receipt.Points = GetPoints(receipt)
	receipt.CalculationDate = time.Now().String()

	var id = uuid.New()
	database.Receipts[id] = receipt

	return id
}

func GetReceiptPoints(id uuid.UUID) (int, bool) {
	var receipt, ok = database.Receipts[id]

	return receipt.Points, ok
}
