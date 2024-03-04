package leetcode_977

import "sort"

// func SortedSquares(nums []int) []int {
// 	n := len(nums)
// 	left := 0
// 	right, index := n-1, n-1

// 	squaredArr := make([]int, n)

// 	for index >= 0 {
// 		leftVal := math.Abs(float64(nums[left]))
// 		rightVal := math.Abs(float64(nums[right]))

// 		if leftVal > rightVal {
// 			squaredArr[index] = int(leftVal) * int(leftVal)
// 			left++
// 		} else {
// 			squaredArr[index] = int(rightVal) * int(rightVal)
// 			right--
// 		}
// 		index--
// 	}

// 	return squaredArr
// }

func SortedSquares(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		num := nums[i] * nums[i]
		nums[i] = num
	}

	sort.Ints(nums)
	return nums
}
