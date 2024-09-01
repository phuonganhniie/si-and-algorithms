package weekly413

import "container/heap"

/*
Description: https://leetcode.com/problems/k-th-nearest-obstacle-queries/description/
*/

// Define a max-heap structure to store distances
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func resultsArray(queries [][]int, k int) []int {
	results := make([]int, len(queries))
	h := &MaxHeap{}
	heap.Init(h)

	for i, query := range queries {
		x, y := query[0], query[1]
		dist := abs(x) + abs(y)
		heap.Push(h, dist)

		if h.Len() > k {
			heap.Pop(h) // remove the largest element if more than k elements
		}

		if h.Len() < k {
			results[i] = -1
		} else {
			results[i] = (*h)[0] // root of the heap is the k-th nearest distance
		}
	}

	return results
}
