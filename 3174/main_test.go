package main

import (
	"testing"
)

func TestClearDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "abc"},    // No digits
		{"cb34", ""},      // All characters removed
		{"a1b2c3", ""},    // Alternating characters and digits
		{"123abc", "abc"}, // Digits at the beginning
		{"abc123", ""},    // Digits at the end
		{"a1b2c3d4", ""},  // Multiple digits and characters
		{"a1b2c3d", "d"},  // Ends with a character
		{"1a2b3c", "c"},   // Digits interleaved with characters
		{"a1b2c", "c"},    // Ends with a digit
		{"a1a1a1", ""},    // Repeated pattern
		{"1a1a1a", "a"},   // Digits at the beginning with repeated pattern
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := clearDigits(test.input)
			if result != test.expected {
				t.Errorf("clearDigits(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}
