package repository

import (
	"github.com/google/uuid"
	"github.com/mostlyworks/FetchReceiptService/models"
)

type Database struct {
	Receipts map[uuid.UUID]models.Receipt
}

func Init() Database {
	database := Database{}

	database.Receipts = make(map[uuid.UUID]models.Receipt)

	return database
}
