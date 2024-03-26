package leetcode_41

import "sort"

func FirstMissingNumber(nums []int) int {
	sort.Ints(nums)

	missingNumber := 1
	for _, num := range nums {
		if num > 0 {
			if num == missingNumber {
				missingNumber++
			} else if num > missingNumber {
				break
			}
		}
	}

	return missingNumber
}

func FirstMissingNumber2(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		// While the current number can be placed in a correct position
		// (not already in place, within the correct range)
		for nums[i] > 0 && nums[i] <= n {
			if nums[nums[i]-1] != nums[i] {
				// Swap nums[i] with nums[nums[i] - 1]
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			} else {
				break
			}
		}
	}

	// Find the first number missing
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// If all numbers are in their correct positions, return n + 1
	return n + 1
}
