package leetcode_1945

func getLucky(s string, k int) int {
	currentSum := 0
	for _, ch := range s {
		position := int(ch - 'a' + 1)
		digits := (position / 10) + (position % 10)
		currentSum += digits
	}

	for i := 1; i < k; i++ {
		digitSum := 0
		for currentSum > 0 {
			digitSum += currentSum % 10
			currentSum /= 10
		}
		currentSum = digitSum
	}

	return currentSum
}
