package biweekly167

func scoreBalance(s string) bool {
	total := 0
	for _, char := range s {
		total += int(char - 'a' + 1)
	}

	leftSum := 0
	for i := 0; i < len(s) - 1; i++ {
		leftSum += int(s[i] - 'a' + 1)
		rightSum := total - leftSum
		if leftSum == rightSum {
			return true
		}
	}
	return false
}