package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	prev := &ListNode{0, head}
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		prev = prev.Next
	}
	prev.Next = slow.Next
	slow = nil

	return head
}
