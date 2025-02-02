package main

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 4, 5, 1, 2},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 1, 3, 4},
			expected: false,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := check(test.nums)
			if result != test.expected {
				t.Errorf("expected %v, but got %v", test.expected, result)
			}
		})
	}
}
