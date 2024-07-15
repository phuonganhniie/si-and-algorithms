package fastandslow

import (
	"github.com/phuonganhniie/educative/helper"
)

func palindrome(head *helper.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	firstHalf := head

	mid := slow

	secondHalf := reverseLinkedList2(mid)
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

	reverseList := reverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reverseList
}

func reverseLinkedList2(head *helper.ListNode) *helper.ListNode {
	var prev = new(helper.ListNode)
	var next = new(helper.ListNode)
	current := head

	for current != nil && next != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
