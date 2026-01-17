package biweekly174

func bestTower(towers [][]int, center []int, radius int) []int {
	cx, cy := center[0], center[1]
	result := []int{-1, -1}
	maxQuality := -1

	for _, tower := range towers {
		x, y, q := tower[0], tower[1], tower[2]

		dist := abs(x-cx) + abs(y-cy)

		if dist <= radius {
			if q > maxQuality || (q == maxQuality && isLexSmaller(x, y, result[0], result[1])) {
				maxQuality = q
				result = []int{x, y}
			}
		}
	}

	return result
}

func isLexSmaller(x1, y1, x2, y2 int) bool {
	if x1 < x2 {
		return true
	}
	if x1 == x2 && y1 < y2 {
		return true
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
