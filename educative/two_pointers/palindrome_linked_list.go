package twopointers

import (
	"github.com/phuonganhniie/leetcode/helper"
)

func palindrome(head *helper.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	firstHalf := head

	mid := slow
	secondHalf := reverseLinkedList(mid)

	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

func reverseLinkedList(head *helper.ListNode) *helper.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := reverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return r
}
