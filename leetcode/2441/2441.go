package leetcode_2441

import "sort"

func FindMaxK(nums []int) int {
	sort.Ints(nums)

	low, high := 0, len(nums)-1

	for low < high {
		if -nums[low] == nums[high] {
			return nums[high]
		}

		if -nums[low] > nums[high] {
			low++
		} else {
			high--
		}
	}
	return -1
}
