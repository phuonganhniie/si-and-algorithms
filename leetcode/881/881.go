package leetcode_881

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	start, end := 0, len(people)-1
	totalBoats := 0

	for start < end {
		totalWeights := people[start] + people[end]
		if totalWeights <= limit {
			start++
			end--
			totalBoats++
		}
		if totalWeights > limit {
			end--
			totalBoats++
		}
	}

	if start == end {
		totalBoats++
	}
	return totalBoats
}
