package binarysearchmodified

func binarySearchRotated(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		// check first half is sorted
		if nums[start] <= nums[mid] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
