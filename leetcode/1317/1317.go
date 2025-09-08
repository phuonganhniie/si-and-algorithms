package leetcode_1317

func isNoZero(num int) bool {
	for num > 0 {
		if num%10 == 0 {
			return false
		}

		num /= 10
	}

	return true
}

func getNoZeroIntegers(n int) []int {
	result := make([]int, 0, n)

	for numA := 1; numA < n; numA++ {
		numB := n - numA

		if isNoZero(numA) && isNoZero(numB) {
			result = append(result, numA, numB)
			break
		}
	}

	return result
}
