package main

import "testing"

func Test3(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "Test 1",
			input:  "abcabcbb",
			output: 3,
		},
		{
			name:   "Test 2",
			input:  "bbbbb",
			output: 1,
		},
		{
			name:   "Test 3",
			input:  "pwwkew",
			output: 3,
		},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if result != test.output {
			t.Errorf("%v failed. Expected %d, but got %d", test.name, test.output, result)
		}
	}
}
