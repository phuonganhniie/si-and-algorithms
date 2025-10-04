package leetcode_11

func maxArea(heights []int) int {
	maxWaterArea := 0
	leftIndex, rightIndex := 0, len(heights)-1

	for leftIndex < rightIndex {
		containerWidth := rightIndex - leftIndex
		limitingHeight := min(heights[leftIndex], heights[rightIndex])
		currentWaterArea := containerWidth * limitingHeight
		maxWaterArea = max(currentWaterArea, maxWaterArea)

		if heights[leftIndex] <= heights[rightIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}
	return maxWaterArea
}
