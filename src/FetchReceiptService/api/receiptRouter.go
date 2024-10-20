package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/mostlyworks/FetchReceiptService/services"
)

var appServer App

func ReceiptRouter(app App) {
	appServer = app
	services.InitPointsService(app.pointsConfig)
	services.InitReceiptService(app.database)
	app.httpServer.HandleFunc("/receipts/process", receiptHandler)
	app.httpServer.HandleFunc("/receipts/{id}/points", receiptPointHandler)
}

func receiptHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handleReceiptPOST(resp, req)
	default:
		http.Error(resp, "Method not Implemented", http.StatusNotImplemented)
	}
}

func receiptPointHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handlePointGET(resp, req)
	default:
		http.Error(resp, "Method not Implemented", http.StatusNotImplemented)
	}
}

func handleReceiptPOST(resp http.ResponseWriter, req *http.Request) {
	// Unmarshal Json from Req body,
	// Calc points and write to DB?
	// Write to "DB"
	// Return "ID"

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(resp, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var receipt models.Receipt

	if err := json.Unmarshal(body, &receipt); err != nil {
		http.Error(resp, "Error parsing request body", http.StatusBadRequest)
		return
	}

	messages := isValidReceipt(receipt)

	if len(messages) > 0 {
		http.Error(resp, strings.Join(messages, "\n"), http.StatusBadRequest)
		return
	}

	var id = services.WriteReciept(receipt)
	var idResp models.Id
	idResp.Id = id.String()

	resp.Header().Set("content-Type", "application/json")
	// This should probably return a 202
	json.NewEncoder(resp).Encode(idResp)
}

func handlePointGET(resp http.ResponseWriter, req *http.Request) {
	// Return Points for "ID"
	var id = req.PathValue("id")

	if !IsValidUUID(id) && len(id) != 0 {
		http.Error(resp, "ID not Valid", http.StatusBadRequest)
		return
	}

	var uuid = uuid.MustParse(id)
	var points, ok = services.GetReceiptPoints(uuid)
	if !ok {
		http.Error(resp, "Receipt Not Found", http.StatusNotFound)
		return
	}

	var pointsResp models.Points
	pointsResp.Points = points
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(pointsResp)
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func isValidReceipt(receipt models.Receipt) []string {
	var messages []string
	// No items
	if !(len(receipt.Items) > 0) {
		messages = append(messages, "No items provided")
	}

	// No date
	if receipt.PurchaseDate.IsZero() {
		messages = append(messages, "Purchase date not provided")
	}

	// No Time
	if receipt.PurchaseTime.IsZero() {
		messages = append(messages, "Purchase time not provided")
	}

	// No Retailer
	if !(len(receipt.Retailer) > 0) {
		messages = append(messages, "Retailer not provided")
	}

	// No Total
	if receipt.Total.IsZero() {
		messages = append(messages, "Total not provided")
	}

	// Check item contents
	if !isValidItems(receipt.Items) {
		messages = append(messages, "Provided items not valid")
	}

	return messages
}

func isValidItems(items []models.Item) bool {
	for _, value := range items {
		if value.Price.IsZero() {
			return false
		}
		if !(len(value.ShortDescription) > 0) {
			return false
		}
	}
	return true
}
