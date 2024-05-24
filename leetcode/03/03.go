package leetcode_03

func LengthOfLongestSubstring(s string) int {
	charactersStorage := make(map[byte]int, len(s))
	longestWindowLen := 0

	start := 0
	for end := 0; end < len(s); end++ {
		currentChar := s[end]

		if _, found := charactersStorage[currentChar]; found {
			start = max(start, charactersStorage[currentChar]+1)
		}

		charactersStorage[currentChar] = end

		currentWindowLen := end - start + 1
		longestWindowLen = max(longestWindowLen, currentWindowLen)
	}
	return longestWindowLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
