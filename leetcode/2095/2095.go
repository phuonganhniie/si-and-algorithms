package leetcode_2095

import "github.com/phuonganhniie/leetcode/helper"

func deleteMiddle(head *helper.ListNode) *helper.ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head
	dummy := &helper.ListNode{Val: 0, Next: head}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		dummy = dummy.Next
	}
	dummy.Next = slow.Next
	slow.Next = nil

	return head
}
