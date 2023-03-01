package main

import (
	"fmt"
	"sort"
)

type InventoryData [][]interface{}

func convertToMap(data [][]interface{}) map[string]int {
	result := make(map[string]int)

	if len(data) < 2 {
		return result
	}

	for _, val := range data {
		itemName := val[1].(string)
		quantity := val[0].(int)
		_, ok := result[itemName]
		if ok {
			result[itemName] += quantity
		} else {
			result[itemName] = quantity
		}
	}
	return result
}

func convertToArray(data map[string]int) InventoryData {
	var result InventoryData
	for key, value := range data {
		result = append(result, []interface{}{value, key})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][1].(string) < result[j][1].(string)
	})
	return result
}

func update(curInv, newInv map[string]int) {
	for key, val := range newInv {
		_, ok := curInv[key]
		if ok {
			curInv[key] += val
		} else {
			curInv[key] = val
		}
	}
}

func updateInventory(array []InventoryData) InventoryData {
	result := make(map[string]int)
	for _, val := range array {
		mapData := convertToMap(val)
		update(result, mapData)
	}

	return convertToArray(result)
}

func main() {
	curInv := InventoryData{
		{21, "Bowling Ball"},
		{2, "Dirty Sock"},
		{1, "Hair Pin"},
		{5, "Microphone"},
	}

	newInv := InventoryData{}

	result := updateInventory([]InventoryData{curInv, newInv})
	fmt.Println(result)
}
