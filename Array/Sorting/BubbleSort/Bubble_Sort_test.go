package main

import (
	"reflect"
	"testing"
)

type Test struct {
	name     string
	input    []int
	expected []int
}

func TestBubbleSort(t *testing.T) {
	tests := []Test{
		{
			name:     "Unsorted array",
			input:    []int{54, 3, 8, 99, 33, 68, 1, 88, 34, 2, 73, 84},
			expected: []int{1, 2, 3, 8, 33, 34, 54, 68, 73, 84, 88, 99},
		},
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element array",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Array with duplicates",
			input:    []int{5, 3, 5, 2, 2, 1},
			expected: []int{1, 2, 2, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
