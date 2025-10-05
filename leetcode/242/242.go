package leetcode_242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, strT := range t {
		m[strT]++
	}

	for _, strS := range s {
		if _, existed := m[strS]; !existed {
			return false
		}
		m[strS]--
		if m[strS] == 0 {
			delete(m, strS)
		}
	}

	return len(m) == 0
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// công thức: character - 'a' = index
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
