package twopointers

func sortColors(colors []int) []int {
	start, current, end := 0, 0, len(colors)-1

	for current <= end {
		switch colors[current] {
		case 0:
			colors[start], colors[current] = colors[current], colors[start]
			current++
			start++
		case 1:
			current++
		case 2:
			colors[current], colors[end] = colors[end], colors[current]
			end--
		}
	}

	return colors
}
