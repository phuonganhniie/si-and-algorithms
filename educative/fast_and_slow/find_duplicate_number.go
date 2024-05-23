package fastandslow

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// part 1: find the insertion point of slow & fast pointers
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// part 2: slow down speed of fast pointer so that it moves with the same speed as slow pointer
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}
