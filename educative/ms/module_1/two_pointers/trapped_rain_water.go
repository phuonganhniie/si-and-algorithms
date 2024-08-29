package twopointers

func rainWater(heights []int) int {
	left, right := 0, len(heights)-1
	leftWallMax, rightWallMax := 0, 0
	storedWater := 0

	for left < right {
		if heights[left] <= heights[right] {
			if heights[left] < leftWallMax {
				storedWater += leftWallMax - heights[left]
			} else {
				leftWallMax = heights[left]
			}
			left++
		}

		if heights[left] > heights[right] {
			if heights[right] < rightWallMax {
				storedWater += rightWallMax - heights[right]
			} else {
				rightWallMax = heights[right]
			}
			right--
		}
	}

	return storedWater
}

func rainWater2(heights []int) int {
	left, right := 0, len(heights)-1
	leftWallMax, rightWallMax := 0, 0
	storedWater := 0

	for left <= right {
		if leftWallMax <= rightWallMax {
			storedWater += max(0, leftWallMax-heights[left])
			leftWallMax = max(leftWallMax, heights[left])
			left++
		} else {
			storedWater += max(0, rightWallMax-heights[right])
			rightWallMax = max(rightWallMax, heights[right])
			right--
		}
	}

	return storedWater
}
