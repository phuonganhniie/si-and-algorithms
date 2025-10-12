package weekly471

func longestBalanced(s string) int {
	n := len(s)
	maxLen := 0

	for i := 0; i < n; i++ {
		freq := make(map[byte]int)
		for j := i; j < n; j++ {
			freq[s[j]]++

			if isBalanced(freq) {
				length := j - i + 1
				maxLen = max(length, maxLen)
			}
		}
	}
	return maxLen
}

func isBalanced(freq map[byte]int) bool {
	var val int
	first := true
	for _, f := range freq {
		if first {
			val = f
			first = false
		} else if f != val {
			return false
		}
	}
	return true
}
