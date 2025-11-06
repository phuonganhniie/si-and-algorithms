package leetcode_33

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// -> mang ben trai da duoc sort: nums[left] to nums[mid]
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// -> mang ben phai da duoc sort: nums[mid] to nums[right]
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if left <= right && nums[left] == nums[right] {
		return left
	}

	return -1
}
