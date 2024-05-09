package helper

type Item struct {
	Value int
	Index int
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].Index = i
	p[j].Index = j
}

func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p)
	item := x.(*Item)
	item.Index = n
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*p = old[0 : n-1]
	return item
}
