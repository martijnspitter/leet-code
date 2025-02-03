package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	output := 1
	desc, inc := 1, 1

	for i := 1; i < len(nums); i++ {
		prev, cur := nums[i-1], nums[i]
		if cur > prev {
			inc++
			desc = 1
		}
		if cur < prev {
			inc = 1
			desc++
		}
		if cur == prev {
			inc = 1
			desc = 1
		}

		if desc > output {
			output = desc
		}
		if inc > output {
			output = inc
		}
	}

	return output
}
