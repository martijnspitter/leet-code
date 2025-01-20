package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		mat      [][]int
		expected int
	}{
		{
			name: "Test 1",
			arr:  []int{1, 3, 4, 2},
			mat: [][]int{
				{1, 4},
				{2, 3},
			},
			expected: 2,
		},
		{
			name: "Test 2",
			arr:  []int{2, 8, 7, 4, 1, 3, 5, 6, 9},
			mat: [][]int{
				{3, 2, 5},
				{1, 4, 6},
				{8, 7, 9},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Main(tt.arr, tt.mat)
			if actual != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}
