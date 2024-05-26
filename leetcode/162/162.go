package leetcode_162

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	leftIsLess := func(mid int) bool {
		if mid == 0 || nums[mid] > nums[mid-1] {
			return true
		}
		return false
	}

	rightIsLess := func(mid int) bool {
		if mid == len(nums)-1 || nums[mid] > nums[mid+1] {
			return true
		}
		return false
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if leftIsLess(mid) && rightIsLess(mid) {
			return mid
		}
		if !leftIsLess(mid) {
			end = mid - 1
			continue
		}
		if !rightIsLess(mid) {
			start = mid + 1
			continue
		}
	}
	return start
}

/**
 * Fixed version
 * Handle Edge Cases: Check if the array has only one element or handle cases where mid is at the boundary (start or end of the array).
 * Update the Binary Search Logic: Add checks to avoid accessing neighbors that don't exist.
 */
func findPeakElementFixed(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2

		leftIsLess := mid == 0 || nums[mid-1] < nums[mid]
		rightIsLess := mid == len(nums)-1 || nums[mid] > nums[mid+1]

		if leftIsLess && rightIsLess {
			return mid
		}

		if mid > 0 && nums[mid-1] > nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
