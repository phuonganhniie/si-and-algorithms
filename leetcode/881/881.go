package leetcode_881

import "sort"

func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	var countBoat int
	l, r := 0, len(people)-1

	for l < r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		countBoat++
	}

	if l == r {
		countBoat++
	}

	return countBoat
}
