package slidingwindow

/*
Visualization how algorithm works
- Initial state:
target = 7, nums = [2, 3, 1, 2, 4, 3]
minLen = 7 (len(nums) + 1)
currentSum = 0
start = 0

- Steps:
end = 0 -> currentSum = 2 [2]
end = 1 -> currentSum = 5 [2, 3]
end = 2 -> currentSum = 6 [2, 3, 1]
end = 3 -> currentSum = 8 [2, 3, 1, 2] -> minLen = 4 -> shrink -> currentSum = 6 [3, 1, 2]
end = 4 -> currentSum = 10 [3, 1, 2, 4] -> minLen = 4 -> shrink -> currentSum = 7 [1, 2, 4] -> minLen = 3 -> shrink -> currentSum = 6 [2, 4]
end = 5 -> currentSum = 9 [2, 4, 3] -> minLen = 3 -> shrink -> currentSum = 7 [4, 3] -> minLen = 2 -> shrink -> currentSum = 3 [3]

- Final state: minLen = 2
*/
func minSubArrayLen(target int, nums []int) int {
	minWindowLen := len(nums) + 1
	currentSum := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		currentSum = currentSum + nums[end]

		for currentSum >= target {
			currentWindowLen := end - start + 1
			minWindowLen = min(minWindowLen, currentWindowLen)
			currentSum = currentSum - nums[start]
			start++
		}
	}

	if minWindowLen == len(nums)+1 {
		return 0
	}
	return minWindowLen
}
