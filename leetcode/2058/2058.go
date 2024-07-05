package leetcode_2058

import (
	"math"

	"github.com/phuonganhniie/leetcode/helper"
)

/*
[Medium] 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points
https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/description/
Created: 2024-07-05
Done   : quên bấm giờ ròi
Attempt: 3
---------------------NOTE---------------------
Time: O(n)
Space: O(n)
*/
func nodesBetweenCriticalPoints(head *helper.ListNode) []int {
	prev := head
	current := head.Next
	criticalPoints := []int{}
	index := 1

	for current != nil && current.Next != nil {
		index++

		if (current.Val > prev.Val && current.Val > current.Next.Val) ||
			(current.Val < prev.Val && current.Val < current.Next.Val) {
			criticalPoints = append(criticalPoints, index)
		}

		prev = current
		current = current.Next
	}

	if len(criticalPoints) < 2 {
		return []int{-1, -1}
	}

	minDistance := math.MaxUint32
	maxDistance := criticalPoints[len(criticalPoints)-1] - criticalPoints[0]

	for i := 1; i < len(criticalPoints); i++ {
		prevNum := criticalPoints[i-1]
		num := criticalPoints[i]
		minDistance = min(minDistance, num-prevNum)
	}
	return []int{minDistance, maxDistance}
}
