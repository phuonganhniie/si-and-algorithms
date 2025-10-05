package leetcode_49

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	anaGroups := [][]string{}
	sortedStringMap := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		key := string(runes)

		sortedStringMap[key] = append(sortedStringMap[key], s)
	}

	for _, arrayStr := range sortedStringMap {
		anaGroups = append(anaGroups, arrayStr)
	}
	return anaGroups
}
