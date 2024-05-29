package leetcode_1877

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)

	maxPairSum := 0
	start, end := 0, len(nums)-1
	for start < end {
		tempSum := nums[start] + nums[end]
		maxPairSum = max(maxPairSum, tempSum)
		start++
		end--
	}
	return maxPairSum
}
