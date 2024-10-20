package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/mostlyworks/FetchReceiptService/models"
	"github.com/mostlyworks/FetchReceiptService/services"
)

var appServer App

func ReceiptRouter(app App) {
	appServer = app
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

	// TODO: Validate Receipt fields.

	receipt.Points = services.GetPoints(receipt)

	var id = uuid.New()

	// Locking might be overkill.
	// Definately kills concurency
	// appServer.database.ReceiptsMu.Lock()
	// defer appServer.database.ReceiptsMu.Unlock()

	appServer.database.Receipts[id] = receipt
	var idResp models.Id
	idResp.Id = id.String()

	resp.Header().Set("content-Type", "application/json")
	// This should probably return a 202
	json.NewEncoder(resp).Encode(idResp)
	//{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
}

func handlePointGET(resp http.ResponseWriter, req *http.Request) {
	// Return Points for "ID"
	var id = req.PathValue("id")
	var uuid = uuid.MustParse(id)
	var receipt, ok = appServer.database.Receipts[uuid]
	if !ok {
		http.Error(resp, "Not Found", http.StatusNotFound)
		return
	}
	var points = receipt.Points
	var pointsResp models.Points
	pointsResp.Points = points
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(pointsResp)
	// { "points": 32 }
}
