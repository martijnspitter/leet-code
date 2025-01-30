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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	list := &ListNode{}
	cur := list

	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}

		sum, newCarry := getSum(val1, val2, carry)
		carry = newCarry

		cur.Next = &ListNode{
			Val: sum,
		}

		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		cur.Next = &ListNode{
			Val: 1,
		}
	}

	return list.Next
}

func getSum(val1, val2, carry int) (sum, newCarry int) {
	total := val1 + val2 + carry
	return total % 10, total / 10
}
