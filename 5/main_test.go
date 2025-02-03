package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "babad",
			expected: "bab", // "aba" is also a valid answer
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Two characters, same",
			input:    "aa",
			expected: "aa",
		},
		{
			name:     "Two characters, different",
			input:    "ab",
			expected: "a", // "b" is also a valid answer
		},
		{
			name:     "Palindrome in the middle",
			input:    "abcba",
			expected: "abcba",
		},
		{
			name:     "Palindrome at the start",
			input:    "abccba",
			expected: "abccba",
		},
		{
			name:     "Palindrome at the end",
			input:    "xyzzyx",
			expected: "xyzzyx",
		},
		{
			name:     "No palindrome longer than 1",
			input:    "abcdefg",
			expected: "a", // Any single character is a valid answer
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("%v failed. Expected %s, got %s", tt.name, tt.expected, result)
			}
		})
	}
}
