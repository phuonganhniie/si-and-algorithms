package leetcode_125

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	newStr := strings.ToLower(strings.ReplaceAll(s, " ", ""))

	start, end := 0, len(newStr)-1

	for start < end {
		if newStr[start] != newStr[end] {
			return false
		}
		start++
		end--
	}

	return true
}
