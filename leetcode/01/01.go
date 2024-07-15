package leetcode

import "sort"

/*
[Medium] 01. Two Sum
https://leetcode.com/problems/two-sum/description/
Created: 2024-07-13
Done   :
Attempt: 3
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: HashMap
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if ordinalNumIdx, exist := m[complement]; exist {
			return []int{ordinalNumIdx, i}
		}
		m[num] = i
	}
	return []int{}
}

/*
---------------------NOTE---------------------
Time: O(nlogn)
Space: O(n)
Approach: Sort + 2 Pointers
*/
func twoSum2(nums []int, target int) []int {
	type numWithIndex struct {
		num   int
		index int
	}

	numsWithIndices := make([]numWithIndex, len(nums))
	for i, num := range nums {
		numsWithIndices[i] = numWithIndex{num, i}
	}

	sort.Slice(numsWithIndices, func(i, j int) bool {
		return numsWithIndices[i].num < numsWithIndices[j].num
	})

	start, end := 0, len(nums)-1
	for start < end {
		sum := numsWithIndices[start].num + numsWithIndices[end].num
		if sum == target {
			return []int{numsWithIndices[start].index, numsWithIndices[end].index}
		}
		if sum < target {
			start++
			continue
		}
		if sum > target {
			end--
			continue
		}
	}
	return []int{}
}
