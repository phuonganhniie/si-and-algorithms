package leetcode_344

func reverseString(s []byte) string {
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return string(s)
}
