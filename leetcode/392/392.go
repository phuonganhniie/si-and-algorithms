package leetcode_392

func IsSubsequence(s string, t string) bool {
	l1, l2 := 0, 0

	for l1 < len(s) && l2 < len(t) {
		if s[l1] == t[l2] {
			l1++
		}
		l2++
	}

	if l1 == len(s) {
		return true
	}
	return false
}
