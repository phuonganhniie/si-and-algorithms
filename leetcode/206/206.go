package leetcode_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := ReverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return r
}
