package leetcode_1493

func longestSubarray(nums []int) int {
	zeroCount := 0
	longestWindow := 0

	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount = zeroCount - 1
			}
			start++
		}

		longestWindow = max(longestWindow, end-start)
	}
	return longestWindow
}
