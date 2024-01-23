package leetcode

// First leetcode O(n^2): DP
func LengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	lis := make([]int, n)
	for i := range lis {
		lis[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}

	maxLen := lis[0]
	for i := 1; i < n; i++ {
		if lis[i] > maxLen {
			maxLen = lis[i]
		}
	}
	return maxLen
}

// Second leetcode O(n log n): Binary Search
func LengthOfLISBS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// tailValues is used to store the minimum tail of all increasing subsequences
	// with length i+1 in tailValues[i]
	lis := make([]int, n)
	size := 0

	lis[0] = nums[0]
	size = 1

	// Iterating from the second element to the end of the array
	for i := 1; i < n; i++ {
		if nums[i] < lis[0] {
			lis[0] = nums[i]
			continue
		}

		if nums[i] > lis[size-1] {
			lis[size] = nums[i]
			size++
			continue
		}

		// nums[i] wants to be the current end candidate of an existing subsequence
		// It will replace ceil value in tailValues
		index := binarySearch(lis, -1, size-1, nums[i])
		lis[index] = nums[i]
	}

	return size
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Binary search to find the smallest number in tailValues
// which is greater than or equal to nums[i]
// If no such number is found, returns -1
func binarySearch(lis []int, left, right, element int) int {
	for right-left > 1 {
		mid := left + (right-left)/2
		if lis[mid] < element {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
