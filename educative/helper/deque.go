package helper

type Deque struct {
	Elements []int
}

func (d *Deque) Append(value int) {
	d.Elements = append(d.Elements, value)
}

// AppendLeft adds an element to the front of deque
func (d *Deque) AppendLeft(value int) {
	d.Elements = append([]int{value}, d.Elements...)
}

// Pops removes and returns the element from the back of deque
func (d *Deque) Pop() int {
	if len(d.Elements) == 0 {
		panic("Pop from empty deque")
	}
	value := d.Elements[len(d.Elements)-1]
	d.Elements = d.Elements[:len(d.Elements)-1]
	return value
}

// PopLeft removes and returns the element from the front of deque
func (d *Deque) PopLeft() int {
	if len(d.Elements) == 0 {
		panic("Pop left from empty deque")
	}
	value := d.Elements[0]
	d.Elements = d.Elements[1:]
	return value
}

// Front returns the element at the front of deque without removing it
func (d *Deque) Front() int {
	if len(d.Elements) == 0 {
		panic("Front from empty deque")
	}
	return d.Elements[0]
}

// Back return the last element of deque without removing it
func (d *Deque) Back() int {
	if len(d.Elements) == 0 {
		panic("Back from empty queue")
	}
	return d.Elements[len(d.Elements)-1]
}
