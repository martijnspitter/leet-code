package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

func check(nums []int) bool {
	breakPoints := 0

	if len(nums) < 2 {
		return true
	}

	for i := range nums {
		if nums[i] > nums[(i+1)%len(nums)] {
			breakPoints++
		}

		if breakPoints > 1 {
			return false
		}
	}

	return breakPoints < 2
}
