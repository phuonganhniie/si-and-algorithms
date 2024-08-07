package leetcode_680

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			isPalindromeL := checkPalindrome(s, start+1, end)
			isPalindromeR := checkPalindrome(s, start, end-1)
			return isPalindromeL || isPalindromeR
		} else {
			start++
			end--
		}
	}
	return true
}

func checkPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
