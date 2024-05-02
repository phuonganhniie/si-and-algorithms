package leetcode_704

func BasicSearch(nums []int, target int) int {
	// sort.Ints(nums)

	begin, end := 0, len(nums)-1

	for begin <= end {
		mid := (begin + end) / 2

		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
