package leetcode_605

func CanPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)

	if n == 0 {
		return true
	}

	for i := 0; i <= length-1; i++ {
		if flowerbed[i] == 0 {
			prevSlot, nextSlot := false, false

			if i == 0 || flowerbed[i-1] == 0 {
				prevSlot = true
			}

			if i == length-1 || flowerbed[i+1] == 0 {
				nextSlot = true
			}

			if prevSlot && nextSlot {
				flowerbed[i] = 1
				n--
			}
		}
	}

	return n <= 0
}
