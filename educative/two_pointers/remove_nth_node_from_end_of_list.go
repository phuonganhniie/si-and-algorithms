package twopointers

import "github.com/phuonganhniie/educative/helper"

func removeNthLastNode(head *helper.ListNode, n int) *helper.ListNode {
	left, right := head, head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next

	return head
}
