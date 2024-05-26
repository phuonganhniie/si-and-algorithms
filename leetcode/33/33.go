package leetcode_33

func Search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	return dfs(nums, target, start, end)
}

// {4, 5, 6, 7, 0, 1, 2} - 0 - 4
// {0, 1, 2, 3, 4, 5, 6}
func dfs(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] == target { // base case
		return mid
	}

	// Determine which are the sorted and rotated sub-arrays
	if nums[mid] > nums[end] { // => array decreasing, start to mid is sorted, mid to end is rotated
		if target >= nums[start] && target < nums[mid] {
			end = mid - 1
			return dfs(nums, target, start, end) // continue search in the left sorted side
		} else {
			start = mid + 1
			return dfs(nums, target, start, end) // continue search in the right rotated side
		}
	} else { // => array increasing, mid to end sorted, start to mid is rotated
		if target > nums[mid] && target <= nums[end] {
			start = mid + 1
			return dfs(nums, target, start, end) // continue search in the right sorted side
		} else {
			end = mid - 1
			return dfs(nums, target, start, end) // continue search in the left rotated side
		}
	}
}
