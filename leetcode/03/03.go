package leetcode_03

func LengthOfLongestSubstring(s string) int {
	trackingWindow := make(map[byte]int)
	maxSubstringLen := 0
	start := 0

	for end := 0; end < len(s); end++ {
		currentChar := s[end]

		// If the character is already in the map, move the left pointer to the right of the previous instance of that character
		if _, isFound := trackingWindow[currentChar]; isFound {
			start = max(start, trackingWindow[currentChar]+1)
		}

		// Update the character's position in the map and expand the window to the right
		trackingWindow[currentChar] = end

		// Update the maximum length
		maxSubstringLen = max(maxSubstringLen, end-start+1)
	}
	return maxSubstringLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
