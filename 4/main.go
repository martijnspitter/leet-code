package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		i := (low + high) / 2
		j := (m+n+1)/2 - i

		if i < m && nums2[j-1] > nums1[i] {
			low = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			high = i - 1
		} else {
			var maxOfLeft, minOfRight float64

			if i == 0 {
				maxOfLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxOfLeft = float64(nums1[i-1])
			} else {
				maxOfLeft = float64(max(nums1[i-1], nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxOfLeft
			}

			if i == m {
				minOfRight = float64(nums2[j])
			} else if j == n {
				minOfRight = float64(nums1[i])
			} else {
				minOfRight = float64(min(nums1[i], nums2[j]))
			}

			return (maxOfLeft + minOfRight) / 2.0
		}
	}

	return 0.0
}
