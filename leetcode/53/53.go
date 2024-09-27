package leetcode_53

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := 0

	for _, num := range nums {
		if currentSum < 0 {
			currentSum = 0
		}

		currentSum += num
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
