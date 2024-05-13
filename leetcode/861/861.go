package leetcode_861

func MatrixScore(grid [][]int) int {
	lenRow := len(grid)
	lenCol := len(grid[0])

	for row := 0; row < lenRow; row++ {
		if grid[row][0] == 0 {
			for col := 0; col < lenCol; col++ {
				grid[row][col] ^= 1 // XOR with 1 will toggle the bit
			}
		}
	}

	score := 0
	for col := 0; col < lenCol; col++ {
		countOnes := 0
		for row := 0; row < lenRow; row++ {
			// Count how many 1s are there in the current column
			countOnes += grid[row][col]
		}
		// Choose the max between countOnes and the number of 0s (which is rows - countOnes)
		score += max(countOnes, lenRow-countOnes) * (1 << (lenCol - col - 1))
	}
	return score
}
