package leetcode_15

import "sort"

/*
[Medium] 15. 3Sum
https://leetcode.com/problems/3sum/description/
Created: 2024-07-11
Done   : 26 mins 22s
Attempt: 3
---------------------NOTE---------------------
Time: O(n^2)
Space: O(1)
Approach: Three pointers + sort + skip duplicates
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// Alternative approach using set to handle duplicates
// Time: O(n^2), Space: O(n) for the set
func threeSumWithSet(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	// Use a map to act as a set for unique triplets
	seen := make(map[[3]int]bool)
	result := [][]int{}

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// Create a key for the set
				triplet := [3]int{nums[i], nums[left], nums[right]}
				if !seen[triplet] {
					seen[triplet] = true
					result = append(result, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
