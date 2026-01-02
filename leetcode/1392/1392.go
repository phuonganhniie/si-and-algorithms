package leetcode_1392

func longestPrefixBruteForce(s string) string {
	n := len(s)

	// thử từ độ dài lớn nhất có thể xuống 1
	for length := n - 1; length >= 1; length-- {
		// lấy prefix và suffix có độ dài length
		prefix := s[0:length]
		suffix := s[n-length:]

		if prefix != suffix {
			continue
		} else {
			return prefix
		}
	}

	return ""
}

func longestPrefix(s string) string {
	n := len(s)

	prefixSuffixMatch := make([]int, n)

	prefixIndex := 0
	suffixIndex := 1

	for suffixIndex < n {
		charInPrefix := s[prefixIndex]
		charInSuffix := s[suffixIndex]

		if charInPrefix == charInSuffix {
			prefixIndex++
			prefixSuffixMatch[suffixIndex] = prefixIndex
			suffixIndex++
		} else {
			if prefixIndex > 0 {
				prefixIndex = prefixSuffixMatch[prefixIndex-1]
			} else {
				prefixSuffixMatch[suffixIndex] = 0
				suffixIndex++
			}
		}
	}

	longestPrefixLength := prefixSuffixMatch[n-1]
	return s[0:longestPrefixLength]
}
