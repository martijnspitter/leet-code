package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	current := ""
	for _, char := range strings.Split(s, "") {
		if strings.Contains(current, char) {
			if len(current) > max {
				max = len(current)
			}
			current = char
		} else {
			current = current + char
		}
	}
	return max
}
