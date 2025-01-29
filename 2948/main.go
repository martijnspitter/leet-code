package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World")
}

type Nums struct {
	nums  []int
	limit int
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	// Create pairs of (value, index)
	pairs := make([][2]int, n)
	for i, num := range nums {
		pairs[i] = [2]int{num, i}
	}

	// Sort pairs by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	// Initialize result array
	result := make([]int, n)

	// Process numbers in groups where each number can be connected to others
	start := 0
	for start < n {
		end := start
		// Find all numbers that can be connected through the limit
		for end+1 < n && pairs[end+1][0]-pairs[end][0] <= limit {
			end++
		}

		// Get all values and indices in this group
		values := make([]int, end-start+1)
		indices := make([]int, end-start+1)
		for i := 0; i < end-start+1; i++ {
			values[i] = pairs[start+i][0]
			indices[i] = pairs[start+i][1]
		}
		sort.Ints(indices)

		// Assign values to the sorted positions
		for i := 0; i < len(indices); i++ {
			result[indices[i]] = values[i]
		}

		start = end + 1
	}

	return result
}
