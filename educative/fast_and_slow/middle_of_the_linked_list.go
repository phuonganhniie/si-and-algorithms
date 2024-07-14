package fastandslow

import "github.com/phuonganhniie/educative/helper"

func getMiddleNode(head *helper.ListNode) *helper.ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
