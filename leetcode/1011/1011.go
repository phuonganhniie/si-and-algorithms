package leetcode_1011

import (
	"slices"
)

func ShipWithinDays(weights []int, days int) int {
	low := slices.Max(weights)
	high := 0
	for _, w := range weights {
		high += w
	}

	canShipInDays := func(capacity int) bool {
		shipDays := 1
		currentWeight := 0

		for _, weight := range weights {
			if currentWeight+weight > capacity {
				shipDays++
				currentWeight = weight

				if shipDays > days {
					return false
				}
			} else {
				currentWeight += weight
			}
		}
		return true
	}

	for low < high {
		mid := low + (high-low)/2

		if canShipInDays(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
