package leetcode_15

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for start := 0; start < len(nums)-2; start++ {
		if start > 0 && nums[start] == nums[start-1] {
			continue
		}

		mid, end := start+1, len(nums)-1
		for mid < end {
			sum := nums[start] + nums[mid] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[mid], nums[end]})

				// skip duplicate elements
				for mid < end && nums[mid] == nums[mid+1] {
					mid++
				}
				for mid < end && nums[end] == nums[end-1] {
					end--
				}
				mid++
				end--
				continue
			}

			if sum < 0 {
				mid++
				continue
			}
			end--
		}
	}
	return result
}
