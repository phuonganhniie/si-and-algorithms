package leetcode_345

func ReverseVowels(s string) string {
	vowelsMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	runeArr := []rune(s)

	l, r := 0, len(runeArr)-1
	for l <= r {
		foundVowels := vowelsMap[runeArr[l]]
		if !foundVowels {
			l++
		}

		foundVowelss := vowelsMap[runeArr[r]]
		if !foundVowelss {
			r--
		}

		if foundVowels && foundVowelss {
			runeArr[l], runeArr[r] = runeArr[r], runeArr[l]
			l++
			r--
		}
	}
	return string(runeArr)
}
