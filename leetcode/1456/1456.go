package leetcode_1456

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'u': true,
		'i': true,
		'e': true,
		'o': true,
	}
	maxVowelCount := 0
	currentVowelCount := 0

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currentVowelCount++
		}
	}
	maxVowelCount = currentVowelCount

	for end := k; end < len(s); end++ {
		if vowels[s[end-k]] {
			currentVowelCount--
		}

		if vowels[s[end]] {
			currentVowelCount++
		}
		maxVowelCount = max(maxVowelCount, currentVowelCount)
	}
	return maxVowelCount
}
