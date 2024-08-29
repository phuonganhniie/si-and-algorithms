package leetcode

func searchRange(nums []int, target int) []int {
	result := make([]int, 2)

	result[0] = findFirstPosition(nums, target)

	result[1] = findSecondPosition(nums, target)

	return result
}

func findFirstPosition(nums []int, target int) int {
	firstPos := -1
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			firstPos = mid
			end = mid - 1
			continue
		}

		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return firstPos
}

func findSecondPosition(nums []int, target int) int {
	secondPos := -1
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			secondPos = mid
			start = mid + 1
			continue
		}

		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return secondPos
}
