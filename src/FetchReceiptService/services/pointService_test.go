package services

import (
	"testing"

	"github.com/mostlyworks/FetchReceiptService/config"
	"github.com/mostlyworks/FetchReceiptService/models"
)

func setup() {
	InitPointsService(config.LoadPointConfig())
}

func TestReceiptTotalPoints(t *testing.T) {
	setup()
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Round dollar & .25", "5.00", pointConfig.TotalRoundedPoints + pointConfig.TotalMutiplePoints},
		{"No points", "35.40", 0},
		{".25", "1.25", pointConfig.TotalMutiplePoints},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := receiptTotalPoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.input, test.output, result)
			}
		})
	}
}

func TestRetailerPoints(t *testing.T) {
	setup()
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Name with Spaces", "T J M A X", 5 * pointConfig.RetailerNamePointMutiplier},
		{"Normal name", "Target", 6 * pointConfig.RetailerNamePointMutiplier},
		{"Latin accented character", "Wàl-mãrt", 5 * pointConfig.RetailerNamePointMutiplier},
		{"Appostrophy", "bj's wholesale", 12 * pointConfig.RetailerNamePointMutiplier},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := retailerPoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.input, test.output, result)
			}
		})
	}
}

// This test will fail if not loading default config (hardcoded assertion)
func TestItemPoints(t *testing.T) {
	setup()
	tests := []struct {
		name   string
		input  []models.Item
		output int
	}{
		{"Mutiple of 3", []models.Item{{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}}, 3},
		{"4 Items",
			[]models.Item{{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"}},
			pointConfig.ItemCountPoints * 2},
		{"5 Items",
			[]models.Item{{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"},
				{ShortDescription: "Gatorade", Price: "2.00"}},
			pointConfig.ItemCountPoints * 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := itemPoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.input, test.output, result)
			}
		})
	}
}

func TestDatePoints(t *testing.T) {
	setup()
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Odd Date", "2024-10-30", 0},
		{"Even Date", "2024-10-31", pointConfig.PurchaseDatePoints},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := datePoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.input, test.output, result)
			}
		})
	}
}

func TestTimePoints(t *testing.T) {
	setup()
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Pre check range", "10:00", 0},
		{"Post Check range", "20:00", 0},
		{"Check Range", "14:33", pointConfig.PurchaseTimePoints},
		{"Check Range", "15:00", pointConfig.PurchaseTimePoints},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := timePoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.input, test.output, result)
			}
		})
	}
}

// This test will fail if not using default config (hardcoded assertion)
func TestGetPoints(t *testing.T) {
	setup()
	tests := []struct {
		name   string
		input  models.Receipt
		output int
	}{
		{"Ex 1", models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []models.Item{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		}, 28},
		{"Ex 2", models.Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []models.Item{
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				}, {
					ShortDescription: "Gatorade",
					Price:            "2.25",
				}, {
					ShortDescription: "Gatorade",
					Price:            "2.25",
				}, {
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
			},
			Total: "9.00",
		}, 109},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetPoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.name, test.output, result)
			}
		})
	}
}
