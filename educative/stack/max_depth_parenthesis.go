package stack

import "github.com/phuonganhniie/educative/helper"

func maxDepthParenthesis(expn string) int {
	var depth, maxDepth int
	stk := helper.Stack{}

	for _, char := range expn {
		switch char {
		case '(':
			stk.Push(int(char))
			depth++

		case ')':
			stk.Pop()
			depth--
		}
		maxDepth = max(maxDepth, depth)
	}
	return maxDepth
}
