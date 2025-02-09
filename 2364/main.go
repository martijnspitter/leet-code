package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func countBadPairs(nums []int) int64 {
	i, j := 0, 1

	var result int64

	for i < j && i < len(nums) {
		j = 1

		for j < len(nums) {
			if i < j && j-i != nums[j]-nums[i] {
				result++
			}
			j++
		}

		i++
	}
	return result
}
