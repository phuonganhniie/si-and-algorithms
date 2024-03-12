package leetcode_1171

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveZeroSumSublists(head *ListNode) *ListNode {
	fakeNode := &ListNode{0, head}
	prefixSum := 0

	zeroSeqMap := make(map[int]*ListNode)
	zeroSeqMap[0] = fakeNode

	for current := fakeNode; current != nil; current = current.Next {
		prefixSum = prefixSum + current.Val
		zeroSeqMap[prefixSum] = current
	}

	prefixSum = 0 // reset
	for current := fakeNode; current != nil; current = current.Next {
		prefixSum = prefixSum + current.Val
		current.Next = zeroSeqMap[prefixSum].Next
	}

	return fakeNode.Next
}
