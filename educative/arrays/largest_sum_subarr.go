package arrays

func MaxSubArraySum(data []int) int {
	maxSum := 0
	currentSum := 0

	for i := 0; i < len(data); i++ {
		currentSum += data[i]

		if currentSum < 0 {
			currentSum = 0
		}

		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
