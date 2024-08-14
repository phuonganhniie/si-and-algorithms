package stack

import "github.com/phuonganhniie/educative/helper"

func bottomInsert(stk *helper.Stack, element int) *helper.Stack {
	var temp int

	if stk.IsEmpty() {
		stk.Push(element)
		return stk
	}

	temp = stk.Pop()

	bottomInsert(stk, element)

	stk.Push(temp)

	return stk
}
