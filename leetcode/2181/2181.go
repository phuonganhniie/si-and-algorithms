package leetcode_2181

import "github.com/phuonganhniie/leetcode/helper"

/*
[Medium] 2181. Merge Nodes in Between Zeros
https://leetcode.com/problems/merge-nodes-in-between-zeros/description/
Created: 2024-07-04
Done   : 12 mins 34s
Attempt: 2
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
*/
func mergeNodes(head *helper.ListNode) *helper.ListNode {
	dummy := head.Next
	nextSum := dummy

	for nextSum != nil {
		sum := 0
		for nextSum.Val != 0 {
			sum += nextSum.Val
			nextSum = nextSum.Next
		}

		dummy.Val = sum

		nextSum = nextSum.Next

		dummy.Next = nextSum

		dummy = dummy.Next
	}
	return head.Next
}
