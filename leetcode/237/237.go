package leetcode_237

import (
	"github.com/phuonganhniie/leetcode/helper"
)

func DeleteNode(node *helper.ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
