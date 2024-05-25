package binarysearch

func binarySearchRotated(nums []int, target int) int {
	start, end := 0, len(nums)-1
	return binarySearch(nums, start, end, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > nums[end] {
		if target >= nums[start] && target < nums[mid] { // = nums[start] <= target < nums[mid]
			end = mid - 1
			return binarySearch(nums, start, end, target)
		} else {
			start = mid + 1
			return binarySearch(nums, start, end, target)
		}
	} else {
		if target > nums[mid] && target <= nums[end] { // = nums[mid] < target <= nums[end]
			start = mid + 1
			return binarySearch(nums, start, end, target)
		} else {
			end = mid - 1
			return binarySearch(nums, start, end, target)
		}
	}
}
