package stack

import "github.com/phuonganhniie/educative/helper"

func sortedInsert(stk *helper.Stack, element int) *helper.Stack {
	var temp int
	if element > stk.Top() {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		sortedInsert(stk, element)
		stk.Push(temp)
	}
	return stk
}
