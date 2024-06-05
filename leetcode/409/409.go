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

// By anh @TuanTuna
func longestPalindrome2(s string) int {
	m := make([]int, 26)
	for _, v := range []byte(s) {
		m[v-'a']++
	}
	ans := 0
	for _, v := range m {
		ans = (ans + (v>>1)<<1) | v&1
	}
	return ans
}
