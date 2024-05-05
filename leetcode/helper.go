package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func BuildLinkedListToArray(head *ListNode) []int {
	var arr []int
	for current := head; current != nil; current = current.Next {
		arr = append(arr, current.Val)
	}
	return arr
}

func FindNode(head *ListNode, val int) *ListNode {
	current := head
	for current != nil {
		if current.Val == val {
			return current
		}
		current = current.Next
	}
	return nil
}
