package leetcode_713

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	prod, count, leftPtr := 1, 0, 0

	for rightPtr, num := range nums {
		prod *= num
		for prod >= k {
			prod /= nums[leftPtr]
			leftPtr++
		}
		count += rightPtr - leftPtr + 1
	}

	return count
}
