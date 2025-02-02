package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello World")
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	new := append(nums1, nums2...)
	slices.Sort(new)
	fmt.Println("new", new)

	middle := len(new) / 2 // converts to int > min()
	isEven := len(new)%2 == 0
	num1 := float64(new[middle])
	num2 := float64(new[middle])
	if isEven {
		num1 = float64(new[middle-1])
	}

	return (num1 + num2) / 2
}
