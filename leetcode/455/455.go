package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	child := 0

	sort.Ints(g)
	sort.Ints(s)

	for _, cookie := range s {
		if child < len(g) && cookie >= g[child] {
			child++ // move to next child
		}
	}

	return child
}
