package leetcode_74

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)    // number of rows
	n := len(matrix[0]) // number of cols
	low, high := 0, m*n-1

	for low <= high {
		mid := low + (high-low)/2
		row := mid / n
		col := mid % n
		midValue := matrix[row][col] // convert to 2D coordinates

		if midValue == target {
			return true
		}
		if midValue <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
