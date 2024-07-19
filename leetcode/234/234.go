package leetcode_234

import "github.com/phuonganhniie/leetcode/helper"

func isPalindrome(head *helper.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow
	rvList := reverseList(mid)

	for head != nil && rvList != nil {
		if head.Val != rvList.Val {
			return false
		}
		head = head.Next
		rvList = rvList.Next
	}
	return true
}

func reverseList(head *helper.ListNode) *helper.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := reverseList(head.Next)
	prev := head.Next
	prev.Next = head
	head.Next = nil
	return r
}
