package main

import "testing"

func Test802(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name: "Test 1",
			input: [][]int{
				{1, 2},
				{2, 3},
				{5},
				{0},
				{5},
				{},
				{},
			},
			expected: []int{2, 4, 5, 6},
		},
		{
			name: "Test 2",
			input: [][]int{
				{1, 2, 3, 4},
				{1, 2},
				{3, 5},
				{0, 4},
				{},
			},
			expected: []int{4},
		},
	}

	for _, test := range tests {
		result := eventualSafeNodes(test.input)

		if len(result) != len(test.expected) {
			t.Errorf("Expected slice of size %d but got %d", len(test.expected), len(result))
		}

		for i, val := range result {
			if test.expected[i] != val {
				t.Errorf("Expect %d but got %d", test.expected[i], val)
			}
		}
	}
}
