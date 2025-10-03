package leetcode_3487

// Link: https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/

func maxSum(nums []int) int {
	seen := make(map[int]bool)
	maxNum := nums[0]

	for _, num := range nums {
		if num > 0 {
			seen[num] = true
		}
		maxNum = max(num, maxNum)
	}

	if len(seen) == 0 {
		return maxNum
	}

	sum := 0
	for num := range seen {
		sum += num
	}
	return sum
}
