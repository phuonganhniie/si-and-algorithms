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
	sum := 0

	for end, num := range nums {
		windows := end - start + 1
		sum += num
		cost := num*windows - sum

		for cost > k {
			sum -= nums[start]
			start++
			windows = end - start + 1
			cost = num*windows - sum
		}

		windows = end - start + 1
		maxFreq = max(maxFreq, windows)
	}
	return maxFreq
}
