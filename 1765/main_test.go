package main

import (
	"testing"
)

func Test1765(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "Test 1",
			input: [][]int{
				{0, 1},
				{0, 0},
			},
			expected: [][]int{
				{1, 0},
				{2, 1},
			},
		},
		{
			name: "Test 2",
			input: [][]int{
				{0, 0, 1},
				{1, 0, 0},
				{0, 0, 0},
			},
			expected: [][]int{
				{1, 1, 0},
				{0, 1, 1},
				{1, 2, 2},
			},
		},
	}

	for _, test := range tests {
		result := highestPeak(test.input)
		for i, row := range result {
			for j, cell := range row {
				if cell != test.expected[i][j] {
					t.Errorf("expected %d, but got %d", test.expected[i][j], cell)
				}
			}
		}
	}
}
