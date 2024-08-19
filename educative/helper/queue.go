package helper

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]interface{}
	front int
	back  int
}

func (q *Queue) Enqueue(value interface{}) {
	if q.size >= capacity {
		return
	}

	q.data[q.back] = value

	q.size++

	q.back = (q.back + 1) % capacity

	return
}

func (q *Queue) Dequeue() interface{} {
	if q.size <= 0 {
		return 0
	}

	q.size--

	var value interface{}
	value = q.data[q.front]

	q.front = (q.front + 1) % capacity

	return value
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Length() int {
	return q.size
}
