package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func twoSum(nums []int, target int) []int {
	indices := []int{}
	sum := 0

	for i := range nums {
		if nums[i] > target {
			continue
		}
		if nums[i] == target {
			return []int{i}
		}
		for j := range nums {
			if i == j {
				continue
			}
			if nums[j] > target {
				continue
			}
			if nums[j] == target {
				return []int{j}
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			if nums[i]+nums[j] > target {
				continue
			}
			if nums[i]+nums[j] < target {
				sum = nums[i] + nums[j]
				indices = append(indices, i, j)
				break
			}
			if nums[i]+nums[j]+sum < target {
				sum = sum + nums[i] + nums[j]
				indices = append(indices, i, j)
				break
			}
			if nums[i]+nums[j]+sum == target {
				sum = sum + nums[i] + nums[j]
				indices = append(indices, i, j)
				return indices
			}

		}
	}

	return indices
}
