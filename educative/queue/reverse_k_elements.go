package queue

import "github.com/phuonganhniie/educative/helper"

func reverseKElementInQueue(que *helper.Queue, k int) *helper.Queue {
	stk := helper.Stack{}

	for i := 0; i < k; i++ {
		value := que.Dequeue()
		stk.Push(value.(int))
	}

	res := helper.Queue{}
	for !stk.IsEmpty() {
		res.Enqueue(stk.Pop())
	}

	for !que.IsEmpty() {
		v := que.Dequeue()
		res.Enqueue(v)
	}
	return &res
}
