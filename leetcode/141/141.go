package leetcode_141

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slowPtr := head
	fastPtr := head.Next

	for fastPtr != nil && fastPtr.Next != nil {
		if slowPtr == fastPtr {
			return true
		}
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	return false
}
