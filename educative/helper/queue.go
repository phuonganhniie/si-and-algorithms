package helper

// ====== Queue implemented by Array ======

const capacity = 100

type Queue struct {
	Size  int
	Data  [capacity]interface{}
	Front int
	Back  int
}

func (q *Queue) Enqueue(value interface{}) {
	if q.Size >= capacity {
		return
	}

	q.Data[q.Back] = value

	q.Size++

	q.Back = (q.Back + 1) % capacity

	return
}

func (q *Queue) Dequeue() interface{} {
	if q.Size <= 0 {
		return 0
	}

	q.Size--

	var value interface{}
	value = q.Data[q.Front]

	q.Front = (q.Front + 1) % capacity

	return value
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) Length() int {
	return q.Size
}

// ====== Queue implemented by Linked List ======
type NodeQueue struct {
	value int
	next  *NodeQueue
}

type QueueLinkedList struct {
	head *NodeQueue
	tail *NodeQueue
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	return q.head.value
}

func (q *QueueLinkedList) Add(value int) {
	temp := &NodeQueue{value: value, next: nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) Remove() int {
	if q.IsEmpty() {
		return -1
	}

	value := q.head.value

	q.head = q.head.next
	q.size--

	return value
}

func (q *QueueLinkedList) AddCircular(value int) {
	temp := &NodeQueue{value: value, next: nil}

	if q.head == nil {
		q.head = temp
		q.tail = temp
		q.tail.next = q.head
	} else {
		temp.next = q.head
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) RemoveCircular() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}

	var value int
	value = q.head.value
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.tail.next = q.head
	}
	q.size--
	return value, true
}
