package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

func countBadPairs(nums []int) int64 {
	n := len(nums)
	result := int64(n * (n - 1) / 2)
	countMap := make(map[int]int)
	goodPairs := int64(0)

	for j := range nums {
		diff := nums[j] - j
		if count, ok := countMap[diff]; ok {
			goodPairs += int64(count)
		}

		countMap[diff]++
	}

	return result - goodPairs
}
