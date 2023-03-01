package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateInventory(t *testing.T) {
	type Tests struct {
		input          []InventoryData
		expectedResult InventoryData
	}

	tests := []Tests{{
		input: []InventoryData{{
			{21, "Bowling Ball"}, {2, "Dirty Sock"}, {1, "Hair Pin"}, {5, "Microphone"},
		}, {
			{2, "Hair Pin"}, {3, "Half-Eaten Apple"}, {67, "Bowling Ball"}, {7, "Toothpaste"},
		}},
		expectedResult: InventoryData{
			{88, "Bowling Ball"}, {2, "Dirty Sock"}, {3, "Hair Pin"}, {3, "Half-Eaten Apple"}, {5, "Microphone"},
			{7, "Toothpaste"}},
	}, {
		input: []InventoryData{{
			{21, "Bowling Ball"}, {2, "Dirty Sock"}, {1, "Hair Pin"}, {5, "Microphone"},
		}, {
			{},
		}},
		expectedResult: InventoryData{
			{21, "Bowling Ball"}, {2, "Dirty Sock"}, {1, "Hair Pin"}, {5, "Microphone"},
		},
	}, {
		input: []InventoryData{{
			{},
		}, {
			{2, "Hair Pin"}, {3, "Half-Eaten Apple"}, {67, "Bowling Ball"}, {7, "Toothpaste"},
		}},
		expectedResult: InventoryData{
			{67, "Bowling Ball"}, {2, "Hair Pin"}, {3, "Half-Eaten Apple"}, {7, "Toothpaste"},
		},
	}, {
		input: []InventoryData{{
			{0, "Bowling Ball"}, {0, "Dirty Sock"}, {0, "Hair Pin"}, {0, "Microphone"},
		}, {
			{1, "Hair Pin"}, {1, "Half-Eaten Apple"}, {1, "Bowling Ball"}, {1, "Toothpaste"},
		}},
		expectedResult: InventoryData{
			{1, "Bowling Ball"}, {0, "Dirty Sock"}, {1, "Hair Pin"}, {1, "Half-Eaten Apple"}, {0, "Microphone"},
			{1, "Toothpaste"}},
	}}

	for _, test := range tests {
		result := updateInventory(test.input)
		assert.Equal(t, test.expectedResult, result)
	}
}
