package twopointers

func isPalindrome(inputString string) bool {
	start, end := 0, len(inputString)-1
	for start <= end {
		if inputString[start] != inputString[end] {
			return false
		}
		start++
		end--
	}
	return true
}
