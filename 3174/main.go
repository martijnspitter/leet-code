package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello World")
}

func clearDigits(s string) string {
	var stack []rune

	for _, char := range s {
		if unicode.IsDigit(char) {
			// Remove the closest non-digit character to the left
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if !unicode.IsDigit(top) {
					break
				}
			}
		} else {
			stack = append(stack, char)
		}
	}

	var result strings.Builder
	for _, char := range stack {
		result.WriteRune(char)
	}

	return result.String()
}
