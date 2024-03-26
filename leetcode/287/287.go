package leetcode_287

func FindDuplicate(nums []int) int {
	var slow, fast = nums[0], nums[0]

	// Phase 1: Finding the intersection point in the cycle
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	// Phase 2: Finding the entrance to the cycle
	ptr1 := nums[0]
	ptr2 := slow
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}

	return ptr1
}

func FindDuplicateBS(nums []int) int {
	low, high := 1, len(nums)-1

	for low < high {
		mid := low + (high-low)/2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
