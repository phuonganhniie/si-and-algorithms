package leetcode_977

import "math"

func SortedSquares(nums []int) []int {
	n := len(nums)
	left := 0
	right, index := n-1, n-1

	squaredArr := make([]int, n)

	for index >= 0 {
		leftVal := math.Abs(float64(nums[left]))
		rightVal := math.Abs(float64(nums[right]))

		if leftVal > rightVal {
			squaredArr[index] = int(leftVal) * int(leftVal)
			left++
		} else {
			squaredArr[index] = int(rightVal) * int(rightVal)
			right--
		}
		index--
	}

	return squaredArr
}
