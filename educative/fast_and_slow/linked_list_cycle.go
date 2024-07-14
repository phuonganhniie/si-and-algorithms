package fastandslow

import "github.com/phuonganhniie/educative/helper"

func detectCycle(head *helper.ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
