package weekly_470

func longestSubsequence(nums []int) int {
	allXor := 0
	allNumsZero := true

	for _, num := range nums {
		allXor ^= num
		if num != 0 {
			allNumsZero = false
		}
	}

	if allNumsZero {
		return 0
	}
	if allXor != 0 {
		return len(nums)
	}
	return len(nums) - 1
}
