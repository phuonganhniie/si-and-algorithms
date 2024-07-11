package leetcode_15

import (
	"sort"
)

/*
[Medium] 15. 3Sum
https://leetcode.com/problems/3sum/description/
Created: 2024-07-11
Done   : 26 mins 22s
Attempt: 3
---------------------NOTE---------------------
Time: O(n^2)
Space: O(1)
Approach: Three pointers + sort
*/
func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		start, end := i+1, len(nums)-1
		for start < end {
			tempSum := nums[i] + nums[start] + nums[end]

			if tempSum == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})

				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}

				start++
				end--
			}

			if tempSum < 0 {
				start++
				continue
			}

			if tempSum > 0 {
				end--
				continue
			}
		}
	}
	return result
}
