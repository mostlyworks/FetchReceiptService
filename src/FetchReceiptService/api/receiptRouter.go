package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mostlyworks/FetchReceiptService/models"
)

func ReceiptRouter(mux *http.ServeMux) {

	mux.HandleFunc("/receipts/process", receiptHandler)
	mux.HandleFunc("/receipts/{id}/points", receiptPointHandler)
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

	// Now we'll try to parse the body. This is similar
	// to JSON.parse in JavaScript.
	if err := json.Unmarshal(body, models.Receipt); err != nil {
		http.Error(resp, "Error parsing request body", http.StatusBadRequest)
		return
	}
}

func handlePointGET(resp http.ResponseWriter, req *http.Request) {
	// Return Points for "ID"
}
