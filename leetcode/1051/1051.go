package leetcode_1051

import "sort"

/*
[Easy] 1051. Height Checker
https://leetcode.com/problems/height-checker/description/
Created: 2024-06-10
Done   : 7 mins 19 secs
Attempt: 3
---------------------NOTE---------------------
Approach: Built-in Sorts
Time: O(n log n)
Space: O(n)

Approach: Counting Sort
Time: O(n)
Space: O(1)
*/
func heightChecker(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	tempHeights := []int{}
	for _, num := range heights {
		tempHeights = append(tempHeights, num)
	}
	sort.Ints(tempHeights)

	count := 0
	left, leftTemp := 0, 0
	for left < len(heights) && leftTemp < len(tempHeights) {
		if heights[left] != tempHeights[leftTemp] {
			count++
		}
		left++
		leftTemp++
	}
	return count
}

func heightCheckerCountingSort(heights []int) int {
	tempHeights := [100]int{}
	for _, height := range heights {
		tempHeights[height]++
	}

	result, pointer := 0, 0
	for height, count := range tempHeights {
		for count > 0 {
			if heights[pointer] != height {
				result++
			}
			pointer++
			count--
		}
	}
	return result
}
