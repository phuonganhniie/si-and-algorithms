package leetcode_153

func findMin(nums []int) int {
	left, right := 0, len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2

		if nums[mid] <= nums[right] {
			right = mid
			continue
		}

		if nums[mid] > nums[right] {
			left = mid + 1
		}
	}
	return nums[left]
}