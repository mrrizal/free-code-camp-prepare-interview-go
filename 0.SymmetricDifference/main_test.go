package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSym(t *testing.T) {
	type Tests struct {
		input          []input
		expectedResult []int
	}

	tests := []Tests{
		{
			input:          []input{[]int{1, 2, 3}, []int{2, 3, 4}},
			expectedResult: []int{1, 4},
		}, {
			input:          []input{[]int{1, 2, 3, 3}, []int{5, 2, 1, 4}},
			expectedResult: []int{3, 4, 5},
		},
		{
			input:          []input{[]int{1, 2, 3, 3}, []int{5, 2, 1, 4, 5}},
			expectedResult: []int{3, 4, 5},
		}, {
			input:          []input{[]int{1, 2, 5}, []int{2, 3, 5}, []int{3, 4, 5}},
			expectedResult: []int{1, 4, 5},
		}, {
			input: []input{
				[]int{3, 3, 3, 2, 5},
				[]int{2, 1, 5, 7},
				[]int{3, 4, 6, 6},
				[]int{1, 2, 3},
				[]int{5, 3, 9, 8},
				[]int{1}},
			expectedResult: []int{1, 2, 4, 5, 6, 7, 8, 9},
		}, {
			input: []input{
				[]int{3, 3, 3, 2, 5},
				[]int{2, 1, 5, 7},
				[]int{3, 4, 6, 6},
				[]int{1, 2, 3}},
			expectedResult: []int{2, 3, 4, 6, 7},
		}}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			result := sym(test.input)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
