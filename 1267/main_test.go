package main

import "testing"

func Test1267(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{
			name: "test 1",
			input: [][]int{
				{1, 0},
				{0, 1},
			},
			output: 0,
		},
		{
			name: "test 2",
			input: [][]int{
				{1, 0},
				{1, 1},
			},
			output: 3,
		},
		{
			name: "test 3",
			input: [][]int{
				{1, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			output: 4,
		},
	}

	for _, test := range tests {
		result := countServers(test.input)
		if test.output != result {
			t.Errorf("expected %d, but got %d", test.output, result)
		}
	}
}
