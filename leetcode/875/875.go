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

func MinEatingSpeed2(piles []int, h int) int {
	minPile := 1
	maxPile := 0
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	getTotalHours := func(speed int) int {
		totalHours := 0
		for _, pile := range piles {
			eachHour := int(math.Ceil(float64(pile) / float64(speed)))
			totalHours += eachHour
		}
		return totalHours
	}

	start, end := minPile, maxPile
	for start <= end {
		mid := start + (end-start)/2
		if getTotalHours(mid) <= h {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

func MinEatingSpeedChatGPT(piles []int, h int) int {
	minPile := 1
	maxPile := 0
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	getTotalHours := func(speed int) int {
		totalHours := 0
		for _, pile := range piles {
			totalHours += (pile + speed - 1) / speed // ceiling of pile/speed
		}
		return totalHours
	}

	start, end := minPile, maxPile
	for start <= end {
		mid := start + (end-start)/2
		if getTotalHours(mid) <= h {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
