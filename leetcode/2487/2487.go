package leetcode_2487

import "github.com/phuonganhniie/leetcode"

func RemoveNodes(head *leetcode.ListNode) *leetcode.ListNode {
	head = reverseLinkedList(head)

	maxVal := head.Val
	dummy := head
	for current := head.Next; current != nil; current = current.Next {
		if current.Val >= maxVal {
			maxVal = current.Val
			dummy = current
		}

		if maxVal > current.Val {
			dummy.Next = current.Next
		}
	}

	return reverseLinkedList(head)
}

func reverseLinkedList(head *leetcode.ListNode) *leetcode.ListNode {
	var reverseList *leetcode.ListNode
	if head == nil {
		return head
	}

	current := head
	for current != nil {
		next := current.Next
		current.Next = reverseList
		reverseList = current
		current = next
	}

	return reverseList
}
