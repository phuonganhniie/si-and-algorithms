package leetcode_1122

import "sort"

/*
[Easy] 1122. Relative Sort Array
https://leetcode.com/problems/relative-sort-array/description/
Created: 2024-06-11
Done   : 30 mins 5 seconds
Attempt: 2
---------------------NOTE---------------------
Time: O(k + n log n)
Space: O(k + n)
Approach:
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	trackingMap := make(map[int]int)
	for _, num := range arr2 {
		trackingMap[num] = 0
	}

	remaining := []int{}
	for _, num := range arr1 {
		if _, exist := trackingMap[num]; exist {
			trackingMap[num]++
		} else {
			remaining = append(remaining, num)
		}
	}

	sort.Ints(remaining)

	result := []int{}
	for _, num := range arr2 {
		count := trackingMap[num]
		for count > 0 {
			result = append(result, num)
			count--
		}
	}
	result = append(result, remaining...)
	return result
}

func relativeSortArrayCountingSort(arr1 []int, arr2 []int) []int {
	counting := make([]int, 1001)
	for _, num := range arr1 {
		counting[num]++
	}

	result := []int{}
	for _, num := range arr2 {
		for counting[num] > 0 {
			result = append(result, num)
			counting[num]--
		}
	}

	for num := 0; num < len(counting); num++ {
		for counting[num] > 0 {
			result = append(result, num)
			counting[num]--
		}
	}

	return result
}
