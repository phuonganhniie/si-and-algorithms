package leetcode_2486

func appendCharacters(s string, t string) int {
	ptrs, ptrt := 0, 0
	for ptrs < len(s) && ptrt < len(t) {
		if s[ptrs] == t[ptrt] {
			ptrs++
			ptrt++
			continue
		}
		ptrs++
	}
	minRemain := len(t[ptrt:])
	return minRemain
}
