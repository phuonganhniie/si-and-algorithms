package twopointers

import (
	"regexp"
	"strings"
)

func reverseWords(sentence string) string {
	words := strings.Fields(sentence)

	start, end := 0, len(words)-1
	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
	return strings.Join(words, " ")
}

// ---------------- Solution 2 ----------------
func reverseWords2(sentence string) string {
	reg := regexp.MustCompile(`\s+`)
	sentence = reg.ReplaceAllString(sentence, " ")
	sentence = strings.TrimSpace(sentence)

	sentence = reverseString(sentence, 0, len(sentence)-1)

	for start, end := 0, 0; end <= len(sentence)-1; end++ {
		if end == len(sentence)-1 || sentence[end] == ' ' {
			var endIdx int
			if end == len(sentence)-1 {
				endIdx = end
			} else {
				endIdx = end - 1
			}

			sentence = reverseString(sentence, start, endIdx)
			start = end + 1
		}
	}
	return sentence
}

func reverseString(s string, start, end int) string {
	runes := []rune(s)

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}
