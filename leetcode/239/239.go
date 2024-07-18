package leetcode_239

import "math"

/*
[Hard] 239. Max Sliding Window
https://leetcode.com/problems/sliding-window-maximum/description/
Created: 2024-07-18
Done   :
Attempt:
---------------------NOTE---------------------
Time: O(N * K)
Space: O(N)
Approach: Sliding Window + Queue
*/
func maxSlidingWindowNaive(nums []int, k int) []int {
	result := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		currentWindow := nums[i : i+k]

		maxNum := math.MinInt16
		for _, num := range currentWindow {
			maxNum = max(maxNum, num)
		}
		result = append(result, maxNum)
	}
	return result
}

func maxSlidingWindowOptimize(nums []int, k int) []int {
	rs := make([]int, len(nums)-k+1)
	queue := []int{0}

	for i := 0; i < k; i++ {
		for len(queue) > 0 && nums[i] >= nums[queue[0]] {
			queue = queue[1:]
		}
		queue = append(queue, i)
	}

	rs[0] = nums[queue[0]]

	for i := k; i < len(nums); i++ {
		// trường hợp dequeue 1: number thứ i >= number thứ 0 trong queue
		for len(queue) > 0 && nums[i] >= nums[queue[0]] {
			queue = queue[1:]
		}

		// trường hợp dequeue 2: number thứ i đạt giới hạn của window size k
		if len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}

		queue = append(queue, i)

		rs[i-k+1] = nums[queue[0]]
	}
	return rs
}
