package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	dummyHead := &ListNode{}
	current := dummyHead
	for head != nil {
		if _, ok := numSet[head.Val]; !ok {
			current.Next = head
			current = current.Next
		}

		head = head.Next
	}

	current.Next = nil

	return dummyHead.Next
}
