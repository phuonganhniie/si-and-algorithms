package leetcode_643

func findMaxAverage(nums []int, k int) float64 {
	var currentSum float64
	for i := 0; i < k; i++ {
		currentSum += float64(nums[i])
	}

	var maxAverage float64
	maxAverage = currentSum / float64(k)

	start := 0
	for end := k; end < len(nums); end++ {
		currentSum = currentSum - float64(nums[start]) + float64(nums[end])
		currentAverage := currentSum / float64(k)
		if currentAverage > maxAverage {
			maxAverage = currentAverage
		}
		start = end - k + 1
	}
	return maxAverage
}
