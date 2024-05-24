package leetcode_03

func LengthOfLongestSubstring(s string) int {
	trackingWindow := make(map[byte]int, len(s))
	longest := 0

	start := 0
	for end := 0; end < len(s); end++ {
		currentChar := s[end]

		if _, found := trackingWindow[currentChar]; found {
			start = max(start, trackingWindow[currentChar]+1)
		}

		trackingWindow[currentChar] = end

		currentLen := end - start + 1
		longest = max(longest, currentLen)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
