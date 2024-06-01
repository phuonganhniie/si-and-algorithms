package leetcode_3110

import "math"

func scoreOfString(s string) int {
	distinctChar := make(map[byte]struct{})
	for _, char := range s {
		distinctChar[byte(char)] = struct{}{}
	}

	start, score := 0, 0
	for end := 1; end < len(s); end++ {
		if _, found := distinctChar[s[start]]; found {
			score += (int(math.Abs(float64(s[start]) - float64(s[end]))))
		}
		start++
	}
	return score
}
