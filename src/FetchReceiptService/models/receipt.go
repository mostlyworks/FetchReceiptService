package models

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"` // Try this as decimal later
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"` // try this as Time?
	PurchaseTime string `json:"purchaseTime"` // try this as Time?
	Items        []Item `json:"items"`
	Total        string `json:"total"` // Try this as decimal later
	Points       int
}
