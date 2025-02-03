package main

import "testing"

func TestLongestMonotonicSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 4, 3, 3, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 3, 3, 3},
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 2, 1},
			expected: 3,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "All increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "All decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 5,
		},
		{
			name:     "Mixed increasing and decreasing",
			nums:     []int{1, 3, 2, 4, 3, 5},
			expected: 2,
		},
		{
			name:     "Plateau",
			nums:     []int{1, 2, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "Alternating",
			nums:     []int{1, 3, 2, 4, 2, 5},
			expected: 2,
		},
		{
			name:     "Longer plateau",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestMonotonicSubarray(tt.nums)
			if result != tt.expected {
				t.Errorf("%v failed. Expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}
