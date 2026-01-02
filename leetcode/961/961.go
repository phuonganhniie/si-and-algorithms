package leetcode_961

import "sort"

func repeatedNTimes(nums []int) int {
	sort.Ints(nums)

	for i, num := range nums {
		if num == nums[i+1] {
			return num
		}
	}
	return 0
}

// Pigeonhole Principle
func repeatedNTimes2(nums []int) int {
	// Check pairs: if repeated element appears n times,
	// it must appear close together
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	// Edge case: pattern like [a, b, c, a] where repeated is at ends
	return nums[len(nums)-1]
}
