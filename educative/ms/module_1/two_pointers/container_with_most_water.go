package twopointers

func containerWithMostWater(height []int) int {
	mostWater := 0
	start, end := 0, len(height)-1
	for start < end {
		distance := end - start
		minHeight := min(height[start], height[end])
		currentWater := minHeight * distance // => area = min(height[start], height[end]) * (end - start)

		if height[start] <= height[end] {
			start++
		} else {
			end--
		}

		mostWater = max(mostWater, currentWater)
	}
	return mostWater
}
