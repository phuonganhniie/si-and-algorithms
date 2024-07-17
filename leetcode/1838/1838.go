package leetcode_1838

import "sort"

/*
[Medium] 1838. Frequency of the Most Frequent Element
https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/
Created: 2024-07-16
Done   :
Attempt: 2
---------------------NOTE---------------------
Time: O(NLogN)
Space: O(LogN)
Intuition:
*/
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	maxFreq := 1
	start := 0
	prefixSum := 0

	for end, num := range nums {
		prefixSum += num
		cost := num*(end-start+1) - prefixSum

		for cost > k {
			prefixSum -= nums[start]
			start++
			cost = num*(end-start+1) - prefixSum
		}

		maxFreq = max(maxFreq, (end - start + 1))
	}
	return maxFreq
}
