package biweekly174

func minOperations(nums []int, target []int) int {
	n := len(nums)
	distinctValues := make(map[int]bool)

	for i := 0; i < n; i++ {
		if nums[i] != target[i] {
			distinctValues[nums[i]] = true
		}
	}

	return len(distinctValues)
}
