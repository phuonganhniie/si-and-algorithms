package leetcode

import "strings"

// func NumDifferentIntegers(word string) int {
// 	s := map[string]struct{}{}
// 	n := len(word)

// 	for i := 0; i < n; i++ {
// 		if isDigit(word[i]) {
// 			for i < n && word[i] == '0' {
// 				i++
// 			}

// 			j := i
// 			for j < n && isDigit(word[j]) {
// 				j++
// 			}
// 			s[word[i:j]] = struct{}{}
// 			i = j
// 		}
// 	}
// 	return len(s)
// }

func NumDifferentIntegers(word string) int {
	validMap := map[string]int{}
	s := []byte(word)

	for i, char := range s {
		if isDigit(char) {
			continue
		}
		s[i] = ' '
	}

	wordList := strings.Fields(string(s))
	for _, w := range wordList {
		val := strings.TrimLeft(w, "0")
		if _, ok := validMap[val]; !ok {
			validMap[val] = 1
		}
	}
	return len(validMap)
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
