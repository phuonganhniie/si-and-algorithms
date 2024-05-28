package leetcode_1498

import (
	"sort"
)

const MOD = 1000000007

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	validSubsequences := 0
	start, end := 0, len(nums)-1

	powerOfTwo := make([]int, len(nums))
	powerOfTwo[0] = 1
	for i := 1; i < len(nums); i++ {
		powerOfTwo[i] = (powerOfTwo[i-1] * 2) % MOD
	}

	for start <= end {
		sum := nums[start] + nums[end]
		if sum <= target {
			validSubsequences = (validSubsequences + powerOfTwo[end-start]) % MOD
			start++
		}

		if sum > target {
			end--
		}
	}
	return validSubsequences
}
