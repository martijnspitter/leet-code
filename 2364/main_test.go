package main

import (
	"fmt"
	"testing"
)

func TestCountBadPairs(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int64
	}{
		{nums: []int{4, 1, 3, 3}, expected: 5},
		{nums: []int{1, 2, 3, 4, 5}, expected: 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("nums=%v", test.nums), func(t *testing.T) {
			result := countBadPairs(test.nums)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
