package leetcode_2816

import "github.com/phuonganhniie/leetcode"

func DoubleIt(head *leetcode.ListNode) *leetcode.ListNode {
	reverseList := leetcode.ReverseLinkedList(head)

	prev := &leetcode.ListNode{}
	current := reverseList
	carry := 0

	for current != nil {
		newVal := current.Val*2 + carry
		current.Val = newVal % 10

		if newVal > 9 {
			carry = 1
		} else {
			carry = 0
		}

		prev = current
		current = current.Next
	}

	if carry != 0 {
		extraNode := &leetcode.ListNode{Val: carry}
		prev.Next = extraNode
	}

	return leetcode.ReverseLinkedList(reverseList)
}

func DoubleItTwoPointers(head *leetcode.ListNode) *leetcode.ListNode {
	head = &leetcode.ListNode{Next: head}

	for current, next := head, head.Next; next != nil; current, next = next, next.Next {
		next.Val *= 2
		current.Val += next.Val / 10
		next.Val %= 10
	}

	if head.Val == 0 {
		head = head.Next
	}

	return head
}
