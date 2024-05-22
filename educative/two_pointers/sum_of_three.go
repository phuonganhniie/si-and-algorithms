package twopointers

import "sort"

func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)

	n := len(nums)

	start, end := 1, n-1

	for i := 0; i < n-2; i++ {
		for start < end {
			tempSum := nums[i] + nums[start] + nums[end]

			if tempSum == target {
				return true
			}
			if tempSum > target {
				end--
				continue
			}
			if tempSum < target {
				start++
				continue
			}
		}
	}
	return false
}
