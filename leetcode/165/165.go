package leetcode_165

import (
	"strconv"
	"strings"
)

func CompareVersion(version1, version2 string) int {
	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")

	maxLen := len(ver1)
	if len(ver2) > maxLen {
		maxLen = len(ver2)
	}

	for i := 0; i < maxLen; i++ {
		p1, p2 := 0, 0

		if i < len(ver1) {
			p1, _ = strconv.Atoi(ver1[i])
		}
		if i < len(ver2) {
			p2, _ = strconv.Atoi(ver2[i])
		}

		if p1 > p2 {
			return 1
		}
		if p1 < p2 {
			return -1
		}
	}
	return 0
}
