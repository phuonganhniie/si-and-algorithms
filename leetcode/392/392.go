package leetcode_392

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	ptrS, ptrT := 0, 0
	for ptrS < len(s) && ptrT < len(t) {
		if s[ptrS] == t[ptrT] {
			ptrT++
			ptrS++
			continue
		}
		ptrT++
	}

	return len(s) == ptrS
}
