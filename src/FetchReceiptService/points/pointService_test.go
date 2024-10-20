package main

import (
	"testing"
)

func TestReceiptTotalPoints(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Round dollar & .25", "5.00", totalRoundPoints + totalMutiplePoints},
		{"No points", "35.40", 0},
		{".25", "1.25", totalMutiplePoints},
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
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Name with Spaces", "T J M A X", 5},
		{"Normal name", "Target", 6},
		{"Latin accented character", "Wàl-mãrt", 5},
		{"Appostrophy", "bj's wholesale", 12},
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

func TestItemPoints(t *testing.T) {
	tests := []struct {
		name   string
		input  []Item
		output int
	}{
		{"Mutiple of 3", []Item{{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"}}, 3},
		{"4 Items", []Item{{"Gatorade", "2.00"}, {"Gatorade", "2.00"}, {"Gatorade", "2.00"}, {"Gatorade", "2.00"}}, itemCountPoints * 2},
		{"5 Items", []Item{{"Gatorade", "2.00"}, {"Gatorade", "2.00"}, {"Gatorade", "2.00"}, {"Gatorade", "2.00"}, {"Gatorade", "2.00"}}, itemCountPoints * 2},
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
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Odd Date", "2024-10-30", 0},
		{"Even Date", "2024-10-31", purchaseDatePoints},
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
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Pre check range", "10:00", 0},
		{"Post Check range", "20:00", 0},
		{"Check Range", "14:33", purchaseTimePoints},
		{"Check Range", "15:00", purchaseTimePoints},
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

func TestGetPoints(t *testing.T) {
	tests := []struct {
		name   string
		input  Receipt
		output int
	}{
		{"Ex 1", Receipt{
			"Target",
			"2022-01-01",
			"13:01",
			[]Item{
				{
					"Mountain Dew 12PK",
					"6.49",
				}, {
					"Emils Cheese Pizza",
					"12.25",
				}, {
					"Knorr Creamy Chicken",
					"1.26",
				}, {
					"Doritos Nacho Cheese",
					"3.35",
				}, {
					"   Klarbrunn 12-PK 12 FL OZ  ",
					"12.00",
				},
			},
			"35.35",
		}, 28},
		{"Ex 2", Receipt{
			"M&M Corner Market",
			"2022-03-20",
			"14:33",
			[]Item{
				{
					"Gatorade",
					"2.25",
				}, {
					"Gatorade",
					"2.25",
				}, {
					"Gatorade",
					"2.25",
				}, {
					"Gatorade",
					"2.25",
				},
			},
			"9.00",
		}, 109},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetPoints(test.input)
			if result != test.output {
				t.Errorf("For input %s, expected %d, but got %d", test.input, test.output, result)
			}
		})
	}
}
