package stack

import "github.com/phuonganhniie/educative/helper"

func reverseKElementInStack(stk *helper.Stack, k int) *helper.Stack {
	var q helper.Queue

	for i := 0; i < k; i++ {
		v := stk.Pop()
		q.Enqueue(v)
	}

	for !q.IsEmpty() {
		stk.Push(q.Dequeue().(int))
	}
	return stk
}
