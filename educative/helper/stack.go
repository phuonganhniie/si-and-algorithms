package helper

// ====== Stack implemented by Array ======

type Stack struct {
	Arr []int
}

func (s *Stack) IsEmpty() bool {
	length := len(s.Arr)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.Arr)
	return length
}

func (s *Stack) Push(value int) {
	s.Arr = append(s.Arr, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() == true {
		return 0
	}
	length := len(s.Arr)
	res := s.Arr[length-1]
	s.Arr = s.Arr[:length-1]
	return res
}

func (s *Stack) Top() int {
	if s.IsEmpty() == true {
		return 0
	}
	length := len(s.Arr)
	res := s.Arr[length-1]
	return res
}

// ====== Stack implemented by Linked List ======

type NodeStack struct {
	val  int
	next *NodeStack
}

type StackLinkedList struct {
	head *NodeStack
	size int
}

// Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
	if s.size == 0 {
		return 0
	}
	return s.size
}

/**
 * IsEmpty() function will return true is size of the linked list is
 * equal to zero or false in all other cases.
 */
func (s *StackLinkedList) IsEmpty() bool {
	return s.head == nil
}

/**
 * First, the Peek() function will check if the stack is empty.
 * If not, it will return the peek value of stack i.e., will return the
 * head value of the linked list.
 */
func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, true // Return 0,true if stack is empty
	}
	return s.head.val, false
}

// Push() function will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
	s.head = &NodeStack{val: value, next: s.head}
	s.size++
}

/**
 * In the pop() function, first it will check that the stack is not empty.
 * Then it will pop the value from the linked list and return it.
 */
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, true
	}

	value := s.head.val
	s.head = s.head.next
	s.size--
	return value, false
}
