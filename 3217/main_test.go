package main

import "testing"

func Test3217(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		head     *ListNode
		expected *ListNode
	}{
		{
			name: "Test 1",
			nums: []int{1, 2, 3},
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		{
			name: "Test 2",
			nums: []int{1},
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val:  2,
									Next: nil,
								},
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Test 3",
			nums: []int{5},
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "Test 4",
			nums: []int{9, 2, 5},
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val:  10,
				Next: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := modifiedList(test.nums, test.head)
			if !compareLinkedLists(result, test.expected) {
				t.Errorf("expected %v, but got %v", linkedListToSlice(test.expected), linkedListToSlice(result))
			}
		})
	}
}

// Helper function to compare two linked lists
func compareLinkedLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// Helper function to convert a linked list to a slice (for easier comparison in tests)
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
