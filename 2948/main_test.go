package main

import "testing"

func Test2948(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		limit  int
		result []int
	}{
		{
			name: "Test 1",
			nums: []int{
				1, 5, 3, 9, 8,
			},
			limit: 2,
			result: []int{
				1, 3, 5, 8, 9,
			},
		},
		{
			name: "Test 2",
			nums: []int{
				1, 7, 6, 18, 2, 1,
			},
			limit: 3,
			result: []int{
				1, 6, 7, 18, 2, 1,
			},
		},
		{
			name: "Test 3",
			nums: []int{
				1, 7, 28, 19, 10,
			},
			limit: 3,
			result: []int{
				1, 7, 28, 19, 10,
			},
		},
	}

	for _, test := range tests {
		result := lexicographicallySmallestArray(test.nums, test.limit)

		if len(result) != len(test.result) {
			t.Errorf("Expected slice of length %d, but got lenght %d", len(test.result), len(result))
		}

		for i, val := range result {
			if val != test.result[i] {
				t.Errorf("Expected %d but got %d", test.result[i], val)
			}
		}
	}
}
