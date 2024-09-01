package binarysearchmodified

func longestSubsequenceBS(nums []int) int {
	temp := make([]int, len(nums), len(nums))
	temp[0] = nums[0]
	lis := 0

	for _, target := range nums {
		left, right := 0, lis

		for left != right {
			mid := left + (right-left)/2

			replacement := temp[mid]

			if replacement < target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		temp[left] = target
		lis = max(left+1, lis)
	}

	return lis
}

func longestSubsequenceDP(nums []int) int {
	return -1
}
