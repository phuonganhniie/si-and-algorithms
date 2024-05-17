package arrays

func MaxSubArraySum(data []int) int {
	maxSum := 0
	currentSum := 0
	for i := 0; i < len(data); i++ {
		currentSum = currentSum + data[i]
		if currentSum < 0 {
			currentSum = 0
		}
		if maxSum < currentSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
