package leetcode_392

func isSubsequence(s string, t string) bool {
	ptrStrS := 0
	ptrStrT := 0

	for ptrStrS < len(s) && ptrStrT < len(t) {
		if s[ptrStrS] == t[ptrStrT] {
			ptrStrS++
		}
		ptrStrT++
	}

	if ptrStrS == len(s) {
		return true
	}
	return false
}
