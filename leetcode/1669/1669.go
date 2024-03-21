package leetcode_1669

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, list1}
	start := dummy

	// run start pointer from first node -> 'a' node
	for i := 0; i < a; i++ {
		start = start.Next
	}

	// run end pointer from the stopped of start pointer -> 'b' node
	end := start
	for i := a; i <= b+1; i++ {
		end = end.Next
	}

	// connect the node just before 'a' to the head of list2
	start.Next = list2

	// find the last node of list2
	lastOfList2 := list2
	for lastOfList2.Next != nil {
		lastOfList2 = lastOfList2.Next
	}

	// connect the last node of list2 to the node after 'b'
	lastOfList2.Next = end

	return dummy.Next
}
