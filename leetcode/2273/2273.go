package leetcode_2273

import (
	"slices"
)

func removeAnagrams(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	result := []string{words[0]}

	for i := 1; i < len(words); i++ {
		lastWord := result[len(result)-1]
		currentWord := words[i]

		if !areAnagrams(lastWord, currentWord) {
			result = append(result, currentWord)
		}
	}

	return result
}

func areAnagrams(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	runes1 := []rune(word1)
	runes2 := []rune(word2)
	slices.Sort(runes1)
	slices.Sort(runes2)

	return slices.Equal(runes1, runes2)
}
