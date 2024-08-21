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

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return head
}

func removeNthLastNode2(head *helper.ListNode, n int) *helper.ListNode {
	headRev := helper.ReverseLinkedList(head)

	ptr := headRev
	for i := 1; i < n-1; i++ {
		ptr = ptr.Next
	}

	ptr.Next = ptr.Next.Next

	head = helper.ReverseLinkedList(headRev)
	return head
}
