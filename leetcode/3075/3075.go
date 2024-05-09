package leetcode_3075

import (
	"container/heap"
	"sort"

	"github.com/phuonganhniie/leetcode/helper"
)

func MaximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		if happiness[i] > happiness[j] {
			return true
		}
		return false
	})

	totalSum := 0
	happinessDecre := 0
	for i := 0; i < k; i++ {
		maxVal := max(0, happiness[i]-happinessDecre)
		totalSum += maxVal
		happinessDecre++
	}

	return int64(totalSum)
}

func MaximumHappinessSum2(happiness []int, k int) int64 {
	priorityQ := make(helper.PriorityQueue, len(happiness))
	for i, v := range happiness {
		priorityQ[i] = &helper.Item{
			Value: v,
			Index: i,
		}
	}
	heap.Init(&priorityQ)

	totalSum := 0
	happinessDecre := 0
	for i := 0; i < k; i++ {
		item := heap.Pop(&priorityQ).(*helper.Item)
		if item.Value-happinessDecre > 0 {
			totalSum += item.Value - happinessDecre
		}
		happinessDecre++
	}

	return int64(totalSum)
}
