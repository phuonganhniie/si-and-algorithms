package weekly413

// Description: https://leetcode.com/contest/weekly-contest-413/problems/check-if-two-chessboard-squares-have-the-same-color/

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	calSum := func(coordinate string) int {
		col := int(coordinate[0] - 'a' + 1)
		row := int(coordinate[1] - '0')
		return col + row
	}

	result := calSum(coordinate1)%2 == calSum(coordinate2)%2

	return result
}
