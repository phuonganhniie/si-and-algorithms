package leetcode_162

/**
 * First version
 */
func FindPeakElement1(nums []int) int {
	if len(nums)-1 <= 1 {
		return len(nums) - 1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] <= nums[mid-1] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

/**
 * Fixed version
 * Handle Edge Cases: Check if the array has only one element or handle cases where mid is at the boundary (start or end of the array).
 * Update the Binary Search Logic: Add checks to avoid accessing neighbors that don't exist.
 */
func FindPeakElement(nums []int) int {
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
