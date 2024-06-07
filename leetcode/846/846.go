package leetcode_846

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	freqMap := make(map[int]int)
	for _, card := range hand {
		freqMap[card]++
	}

	for _, card := range hand {
		if freqMap[card] > 0 {
			for i := 0; i < groupSize; i++ {
				consecutiveCard := card + i
				if freqMap[consecutiveCard] > 0 {
					freqMap[consecutiveCard]--
					continue
				}
				return false
			}
		}
	}
	return true
}
