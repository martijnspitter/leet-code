package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		output []int
	}{
		{
			name:   "Test 1",
			input:  []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			name:   "Test 2",
			input:  []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			name:   "Test 3",
			input:  []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for _, test := range tests {
		output := twoSum(test.input, test.target)
		if len(output) != len(test.output) {
			t.Errorf("%s failed. Expected array of length %d but got %d", test.name, len(test.output), len(output))
			t.FailNow()
		}

		sum := 0

		for _, val := range output {
			sum = sum + test.input[val]
		}
		if sum != test.target {
			fmt.Println(output)
			t.Errorf("%s failed. Expected %d but got %d", test.name, test.target, sum)
		}
	}
}
