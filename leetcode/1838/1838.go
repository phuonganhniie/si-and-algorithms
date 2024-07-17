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
	currentTotal := 0
	for end := 0; end < len(nums); end++ {
		currentTotal += nums[end]

		targetTotal := nums[end] * (end - start + 1)
		neededIncrements := targetTotal - currentTotal
		for neededIncrements > k {
			currentTotal -= nums[start]
			start++
			neededIncrements = nums[end]*(end-start+1) - currentTotal
		}

		windowLength := end - start + 1
		maxFreq = max(maxFreq, windowLength)
	}
	return maxFreq
}
