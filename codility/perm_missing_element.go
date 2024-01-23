package codility

func SolutionLesson3_2(A []int) int {
	n := len(A)
	if n == 0 {
		return 1
	}

	expectedSum := (n + 1) * (n + 2) / 2
	realitySum := 0

	for _, num := range A {
		realitySum += num
	}

	missingNum := expectedSum - realitySum
	return missingNum
}
