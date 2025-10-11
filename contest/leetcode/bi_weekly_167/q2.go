package biweekly167

func longestSubarray(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	maxLen, currentLen := 2, 2
	for i := 2; i < n; i++ {
		if nums[i] == nums[i-1]+nums[i-2] {
			currentLen++
		} else {
			currentLen = 2
		}
		maxLen = max(currentLen, maxLen)
	}
	return maxLen
}
