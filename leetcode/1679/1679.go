package leetcode_1679

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	start, end := 0, len(nums)-1

	operations := 0
	for start < end {
		tempSum := nums[start] + nums[end]
		if tempSum == k {
			operations++
			start++
			end--
		}
		if tempSum < k {
			start++
		}
		if tempSum > k {
			end--
		}
	}
	return operations
}
