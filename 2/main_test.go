package main

import (
	"testing"
)

func Test2(t *testing.T) {
	tests := []struct {
		name   string
		input  []*ListNode
		output *ListNode
	}{
		{
			name: "Test 1",
			input: []*ListNode{
				&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
				&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			},
			output: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
		{
			name: "Test 2",
			input: []*ListNode{
				&ListNode{Val: 0, Next: nil},
				&ListNode{Val: 0, Next: nil},
			},
			output: &ListNode{Val: 0, Next: nil},
		},
		{
			name: "Test 3",
			input: []*ListNode{
				&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}},
				&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}},
			},
			output: &ListNode{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addTwoNumbers(tt.input[0], tt.input[1])
			if !compareLinkedList(result, tt.output) {
				t.Errorf("got %v, want %v", result, tt.output)
			}
		})
	}
}

func compareLinkedList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}
