package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1String := traverse(l1)
	l2String := traverse(l2)

	sum := sum(l1String, l2String)

	var list *ListNode

	for _, val := range sum {
		list = &ListNode{
			Val:  val,
			Next: list,
		}

	}

	return list
}

func traverse(l *ListNode) string {
	val := fmt.Sprintf("%d", l.Val)

	if l.Next == nil {
		return val
	}

	return traverse(l.Next) + val
}

func sum(l1String, l2String string) []int {
	l1num, err := strconv.Atoi(l1String)
	if err != nil {
		return nil
	}
	l2num, err := strconv.Atoi(l2String)
	if err != nil {
		return nil
	}

	sum := fmt.Sprintf("%d", l1num+l2num)
	res := []int{}
	strings := strings.Split(sum, "")
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		res = append(res, n)
	}

	return res
}
