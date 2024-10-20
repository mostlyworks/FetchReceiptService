package models

import (
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Item struct {
	ShortDescription string          `json:"shortDescription"`
	Price            decimal.Decimal `json:"price"` // Try this as decimal later
}

type Receipt struct {
	Retailer        string          `json:"retailer"`
	PurchaseDate    Date            `json:"purchaseDate"` // try this as Time?
	PurchaseTime    Time            `json:"purchaseTime"` // try this as Time?
	Items           []Item          `json:"items"`
	Total           decimal.Decimal `json:"total"` // Try this as decimal later
	Points          int
	CalculationDate string
}

type Time struct {
	time.Time
}

type Date struct {
	time.Time
}

const DateLayout = "2006-01-02"
const TimeLayout = "15:04"

func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(TimeLayout, s)
	return
}

func (ct *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(DateLayout, s)
	return
}
