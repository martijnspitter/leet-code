package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0
	chars := make(map[rune]int)

	for right, char := range s {
		if index, ok := chars[char]; ok && index >= left {
			left = index + 1
		}
		chars[char] = right
		length := right - left + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
