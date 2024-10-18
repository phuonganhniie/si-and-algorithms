package leetcode_33

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	return searchHandler(nums, target, start, end)
}

func searchHandler(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > nums[end] {
		if target >= nums[start] && target < nums[mid] {
			end = mid - 1
			return searchHandler(nums, target, start, end)
		} else {
			start = mid + 1
			return searchHandler(nums, target, start, end)
		}
	} else {
		if target > nums[mid] && target <= nums[end] {
			start = mid + 1
			return searchHandler(nums, target, start, end)
		} else {
			end = mid - 1
			return searchHandler(nums, target, start, end)
		}
	}
}
