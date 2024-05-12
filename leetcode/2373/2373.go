package leetcode_2373

func LargestLocal(grid [][]int) [][]int {
	lenGrid := len(grid[0])
	largestLocal := make([][]int, lenGrid-2)
	for i := range largestLocal {
		largestLocal[i] = make([]int, lenGrid-2)
	}

	findLargestElement := func(row, col int) int {
		maxElement := 0
		for simulateRow := row; simulateRow < row+3; simulateRow++ {
			for simulateCol := col; simulateCol < col+3; simulateCol++ {
				comparedElement := grid[simulateRow][simulateCol]
				maxElement = max(maxElement, comparedElement)
			}
		}
		return maxElement
	}

	for row := 0; row < lenGrid-2; row++ {
		for col := 0; col < lenGrid-2; col++ {
			largestLocal[row][col] = findLargestElement(row, col)
		}
	}
	return largestLocal
}
