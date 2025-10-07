package leetcode_128

func longestConsecutive(nums []int) int {
	trackMap := make(map[int]bool)
	for _, num := range nums {
		trackMap[num] = true
	}

	longestSubsLen := 0
	for num := range trackMap {
		currentSubsLen := 1
		nextNum := num + 1
		if !trackMap[num-1] {
			for {
				if trackMap[nextNum] {
					currentSubsLen++
					nextNum++
				} else {
					break
				}
			}
			longestSubsLen = max(currentSubsLen, longestSubsLen)
		}
	}
	return longestSubsLen
}