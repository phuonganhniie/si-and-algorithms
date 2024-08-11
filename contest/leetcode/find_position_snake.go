package leetcode

/*
Description: https://leetcode.com/contest/weekly-contest-410/problems/snake-in-matrix/description/
*/
func finalPositionOfSnake(n int, commands []string) int {
	var r, c int

	for _, command := range commands {
		switch command {
		case "UP":
			r--
		case "DOWN":
			r++
		case "LEFT":
			c--
		case "RIGHT":
			c++
		}
	}
	return r*n + c
}
