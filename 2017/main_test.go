package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "test 1",
			grid: [][]int{
				{2, 5, 4},
				{1, 5, 1},
			},
			expected: 4,
		},
		{
			name: "test 2",
			grid: [][]int{
				{3, 3, 1},
				{8, 5, 2},
			},
			expected: 4,
		},
		{
			name: "test 3",
			grid: [][]int{
				{1, 3, 1, 15},
				{1, 3, 3, 1},
			},
			expected: 7,
		},
	}

	for _, test := range tests {
		output := gridGame(test.grid)
		if output != test.expected {
			t.Errorf("expected %d, but got %d", test.expected, output)
		}
	}
}
