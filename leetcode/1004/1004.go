package leetcode_1004

func longestOnes(nums []int, k int) int {
	flipped, longestWindow := 0, 0

	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			flipped++
		}

		for flipped > k {
			if nums[start] == 0 {
				flipped--
			}
			start++
		}
		longestWindow = max(longestWindow, end-start+1)
	}
	return longestWindow
}
