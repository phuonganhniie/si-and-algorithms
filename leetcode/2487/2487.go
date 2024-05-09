package leetcode_2487

import (
	"github.com/phuonganhniie/leetcode/helper"
)

func RemoveNodes(head *helper.ListNode) *helper.ListNode {
	head = helper.ReverseLinkedList(head)

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

	return helper.ReverseLinkedList(head)
}
