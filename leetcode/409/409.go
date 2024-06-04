package leetcode_409

func longestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	freqCharMap := make(map[rune]int)
	for _, ch := range s {
		freqCharMap[ch]++
	}

	longestPad := 0
	oddFreqFound := false
	for _, count := range freqCharMap {
		if count%2 == 0 {
			longestPad += count
		}
		if count%2 != 0 {
			longestPad += count - 1
			oddFreqFound = true
		}
	}

	if oddFreqFound {
		longestPad += 1
	}
	return longestPad
}
