package leetcode_33

func Search(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	return dfs(nums, target, begin, end)
}

// {4, 5, 6, 7, 0, 1, 2} - 0 - 4
// {0, 1, 2, 3, 4, 5, 6}
func dfs(nums []int, target, begin, end int) int {
	if begin > end {
		return -1
	}

	mid := (begin + end) / 2
	if nums[mid] == target { // base case
		return mid
	}

	// Determine which are the sorted and rotated sub-arrays
	if nums[mid] > nums[end] { // left side is sorted, right is rotated
		if target >= nums[begin] && target < nums[mid] {
			return dfs(nums, target, begin, mid-1) // continue search in the left sorted side
		} else {
			return dfs(nums, target, mid+1, end) // continue search in the right rotated side
		}
	} else { // right side is sorted, left is rotated
		if target > nums[mid] && target <= nums[end] {
			return dfs(nums, target, mid+1, end) // continue search in the right sorted side
		} else {
			return dfs(nums, target, begin, mid-1) // continue search in the left rotated side
		}
	}
}
