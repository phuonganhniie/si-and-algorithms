package stack

import (
	"github.com/phuonganhniie/educative/helper"
)

func sortStack(stk *helper.Stack) *helper.Stack {
	if stk.IsEmpty() {
		return stk
	}

	top := stk.Pop()

	sortStack(stk)

	stk = sortedInsert(stk, top)

	return stk
}
