package leetcode_162

func findPeakElement(nums []int) int {
	gtThanLeft := func(mid int) bool {
		if mid == 0 || nums[mid-1] < nums[mid] {
			return true
		}
		return false
	}

	gtThanRight := func(mid int) bool {
		if mid == len(nums)-1 || nums[mid] > nums[mid+1] {
			return true
		}
		return false
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if gtThanLeft(mid) && gtThanRight(mid) {
			return mid
		}

		if !gtThanLeft(mid) {
			end = mid - 1
			continue
		}

		if !gtThanRight(mid) {
			start = mid + 1
			continue
		}
	}
	return start
}
