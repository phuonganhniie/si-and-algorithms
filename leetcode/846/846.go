package leetcode_846

import "sort"

/*
[Medium] 846. Hand of Straights
https://leetcode.com/problems/hand-of-straights/description/
Created: 2024-06-05
Done   : 25 mins
Attempt: 1
---------------------NOTE---------------------
Time: O(n log n)
Space: O(n)

Approach:
1. Check for divisibility: If the number of cards isn't divisible by groupSize, it's impossible.
2. Sort:  Make it easier to find consecutive groups.
3. Frequency Map: Count occurrences of each card.
4. Greedy Iteration:
+ Start with the smallest card.
+ For each card, check if it and the next (groupSize - 1) consecutive cards exist in the map.
+ If so, decrement their frequencies.
+ If not, return false.
5. Success: If we iterate through all cards without returning false, it's possible to arrange the hand.
*/
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
