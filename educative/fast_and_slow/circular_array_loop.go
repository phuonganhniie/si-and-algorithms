package fastandslow

func circularArrayLoop(nums []int) bool {
	n := len(nums)

	nextIndex := func(currentIndex int) int {
		step := nums[currentIndex]
		nextIndex := (currentIndex + step) % n

		if nextIndex < 0 {
			return nextIndex + n
		}
		return nextIndex
	}

	isSameDirection := func(currentIndex, nextIndex int) bool {
		return nums[currentIndex]*nums[nextIndex] > 0 // xác định hướng đi, true: tiến - false: lùi
	}

	for i := 0; i < n; i++ {
		step := nums[i]

		if step == 0 {
			continue
		}

		slow, fast := i, nextIndex(i)

		for isSameDirection(i, fast) && isSameDirection(i, nextIndex(fast)) {
			if slow == fast {
				if slow == nextIndex(slow) {
					break
				}
				return true
			}

			slow = nextIndex(slow)
			fast = nextIndex(nextIndex(fast))
		}

		// đánh dấu tất cả các chỉ số đã kiểm tra trong vòng lặp này
		for isSameDirection(i, nextIndex(i)) {
			nextIndex := nextIndex(i)
			nums[i] = 0
			i = nextIndex
		}
	}

	return false
}
