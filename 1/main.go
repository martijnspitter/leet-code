package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func twoSum(nums []int, target int) []int {
	indices := map[int]int{}

	for i, val := range nums {
		remainder := target - val

		j, ok := indices[remainder]
		if !ok {
			indices[val] = i
			continue
		}

		return []int{i, j}
	}

	return nil
}
