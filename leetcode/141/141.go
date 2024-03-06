package leetcode_141

/**
 * 141. The Cycle of Linked List
 * https://leetcode.com/problems/linked-list-cycle/submissions/1195412987/?envType=daily-question&envId=2024-03-06
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Two Pointer
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

// Map
func HasCycleMap(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	visitedMap := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := visitedMap[head]; ok {
			return true
		}

		visitedMap[head] = true
		head = head.Next
	}
	return false
}
