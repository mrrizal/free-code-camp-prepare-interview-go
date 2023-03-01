package main

import (
	"sort"
)

type input []int

func in(inputA int, inputB []int) bool {
	for _, val := range inputB {
		if inputA == val {
			return true
		}
	}
	return false
}

func findSym(inputs []input, indexA, indexB int) []int {
	var result []int
	for _, val := range inputs[indexA] {
		if !in(val, inputs[indexB]) && !in(val, result) {
			result = append(result, val)
		}
	}
	return result
}

func sym(inputs []input) []int {
	result := findSym(inputs, 0, 1)
	for _, val := range findSym(inputs, 1, 0) {
		if !in(val, result) {
			result = append(result, val)
		}
	}

	if len(inputs) > 2 {
		for i := 2; i < len(inputs); i++ {
			inputs[i-1] = result
			tempA := findSym(inputs, i-1, i)
			tempB := findSym(inputs, i, i-1)
			result = append(tempA, tempB...)
		}
	}

	sort.Ints(result)
	return result
}

func main() {

}
