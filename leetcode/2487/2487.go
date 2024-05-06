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
	// Base case:
	// If the linked list is empty or has only one node,
	// return the head as it is already reversed.
	if head == nil || head.Next == nil {
		return head
	}

	// Recursive step:
	// Reverse the linked list starting
	// from the second node (head.next).
	newHead := reverseLinkedList(head.Next)

	// Save a reference to the node following
	// the current 'head' node.
	front := head.Next

	// Make the 'front' node point to the current
	// 'head' node in the reversed order.
	front.Next = head

	// Break the link from the current 'head' node
	// to the 'front' node to avoid cycles.
	head.Next = nil

	// Return the 'newHead,' which is the new
	// head of the reversed linked list.
	return newHead
}
