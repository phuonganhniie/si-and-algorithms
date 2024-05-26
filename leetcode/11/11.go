package leetcode_11

func maxArea(height []int) int {
	low, high := 0, len(height)-1
	maxContainerArea := 0

	for low < high {
		width := high - low
		currentHeight := min(height[low], height[high])
		currentArea := width * currentHeight
		maxContainerArea = max(maxContainerArea, currentArea)

		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return maxContainerArea
}
