package fiverr

func GetCableLength(nRows int, nCols int, image []string) int {
	const mod = 1000000007

	colorRowMap := make(map[rune][]int)
	colorColMap := make(map[rune][]int)

	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			color := rune(image[row][col])
			colorRowMap[color] = append(colorRowMap[color], row)
			colorColMap[color] = append(colorColMap[color], col)
		}
	}

	totalCabLength := 0

	for color := range colorRowMap {
		rowCount := make(map[int]int)
		colCount := make(map[int]int)
		for _, row := range colorRowMap[color] {
			rowCount[row]++
		}
		for _, col := range colorColMap[color] {
			colCount[col]++
		}

		for _, count := range rowCount {
			totalCabLength += (count - 1) % mod
			totalCabLength %= mod
		}
		for _, count := range colCount {
			totalCabLength += (count - 1) % mod
			totalCabLength %= mod
		}
	}

	return totalCabLength
}
