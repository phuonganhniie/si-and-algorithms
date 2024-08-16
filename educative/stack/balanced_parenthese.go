package stack

import "github.com/phuonganhniie/educative/helper"

func isBalancedParenthesis(expn string) bool {
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stk helper.Stack
	for _, r := range expn {
		switch r {
		case '{':
			stk.Push('{')
		case '(':
			stk.Push('(')
		case '[':
			stk.Push('[')
		}

		if ch2, find := pairs[r]; find {
			ch1 := stk.Pop()
			if ch1 != int(ch2) {
				return false
			}
		}
	}
	return stk.IsEmpty()
}
