package stack

import "github.com/phuonganhniie/educative/helper"

func reverseStack(stk *helper.Stack) *helper.Stack {
	if stk.IsEmpty() {
		return stk
	}

	element := stk.Pop()
	reverseStack(stk)

	bottomInsert(stk, element)
	return stk
}
