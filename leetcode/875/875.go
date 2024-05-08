package leetcode_875

import (
	"math"
	"sort"
)

func MinEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)

	low := 1
	high := piles[len(piles)-1]

	searchMinSpeed := func(low, high, target int) int {
		for low <= high {
			totalHours := 0
			mid := low + (high-low)/2

			for _, pile := range piles {
				hour := math.Ceil(float64(pile) / float64(mid))
				totalHours += int(hour)
			}

			if totalHours <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		return low
	}

	return searchMinSpeed(low, high, h)
}
