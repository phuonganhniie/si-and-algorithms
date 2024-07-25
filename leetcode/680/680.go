package leetcode_680

func validPalindrome(s string) bool {
	slow, high := 0, len(s)-1

	for slow < high {
		if s[slow] == s[high] {
			slow++
			high--
			continue
		}
		return isPalindrome(s, slow+1, high) || isPalindrome(s, slow, high-1)
	}
	return true
}

func isPalindrome(s string, slow, high int) bool {
	for slow < high {
		if s[slow] != s[high] {
			return false
		}
		slow++
		high--
	}
	return true
}
