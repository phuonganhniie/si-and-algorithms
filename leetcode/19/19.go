package leetcode_19

import "github.com/phuonganhniie/leetcode/helper"

/*
[Medium] 19. Remove Nth Last Node
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Created: 2024-07-10
Done   :
Attempt: 1
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: Linked List with 2 Pointers
*/
func removeNthFromEnd(head *helper.ListNode, n int) *helper.ListNode {
	left, right := head, head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return head
}
