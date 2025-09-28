package leetcode_976

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		a, b, c := nums[i-2], nums[i-1], nums[i]
		if a+b > c {
			return a + b + c
		}
	}
	return 0
}
