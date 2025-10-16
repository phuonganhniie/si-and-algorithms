package leetcode_643

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxAvg := float64(0)
	for i := 0; i < k; i++ {
		sum += nums[i]
		maxAvg = float64(sum) / float64(k)
	}

	for i := 1; i <= len(nums)-k; i++ {
		start := i - 1
		end := i + k - 1
		sum = sum - nums[start] + nums[end]
		maxAvg = max(float64(sum)/float64(k), float64(maxAvg))
	}
	return maxAvg
}
