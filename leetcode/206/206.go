package leetcode_206

import "github.com/phuonganhniie/leetcode/helper"

func ReverseList(head *helper.ListNode) *helper.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1. next node luôn đằng sau current node
	// 2. trỏ current node đến node nữa nữa của next node
	// 3. trỏ next node vào head
	// 4. gán head vào next node
	currentNode := head
	for currentNode.Next != nil {
		nextNode := currentNode.Next
		currentNode.Next = nextNode.Next
		nextNode.Next = head
		head = nextNode
	}
	return head
}
