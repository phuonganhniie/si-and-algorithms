package helper

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

func ReverseLinkedList(head *ListNode) *ListNode {
	// Base case:
	// If the linked list is empty or has only one node,
	// return the head as it is already reversed.
	if head == nil || head.Next == nil {
		return head
	}

	// Recursive step:
	// Reverse the linked list starting
	// from the second node (head.next).
	newHead := ReverseLinkedList(head.Next)

	// Save a reference to the node following
	// the current 'head' node.
	prev := head.Next

	// Make the 'prev' node point to the current
	// 'head' node in the reversed order.
	prev.Next = head

	// Break the link from the current 'head' node
	// to the 'prev' node to avoid cycles.
	head.Next = nil

	// Return the 'newHead,' which is the new
	// head of the reversed linked list.
	return newHead
}
