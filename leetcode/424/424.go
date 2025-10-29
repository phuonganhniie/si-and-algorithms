package leetcode_424

func characterReplacement(s string, k int) int {
	count := [26]int{}
	start, maxRepeat, maxLen := 0, 0, 0

	for end := 0; end < len(s); end++ {
		count[s[end]-'A']++

		maxRepeat = max(maxRepeat, count[s[end]-'A'])

		if end-start+1-maxRepeat > k {
			count[s[start]-'A']--
			start++
		}

		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}
