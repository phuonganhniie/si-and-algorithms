package leetcode_1208

import "math"

// Binary Search + Fixed Sliding Window approach
func equalSubstringBinarySearch(s string, t string, maxCost int) int {
	cost := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// convert byte to ASCII
		charS := int(s[i]) % 256
		charT := int(t[i]) % 256
		cost[i] = int(math.Abs(float64(charS - charT)))
	}

	isValidCostLen := func(mid int) bool {
		currentCost := 0
		for start := 0; start < mid; start++ {
			currentCost += cost[start]
		}
		if currentCost <= maxCost {
			return true
		}

		for end := mid; end < len(cost); end++ {
			windowStartIndx := end - mid
			currentCost = currentCost + cost[end] - cost[windowStartIndx]
			if currentCost <= maxCost {
				return true
			}
		}
		return false
	}

	maxLen := 0
	start, end := 0, len(s)
	for start <= end {
		mid := start + (end-start)/2
		if isValidCostLen(mid) {
			maxLen = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return maxLen
}

// Sliding Window approach
func equalSubstringSlidingWindow(s string, t string, maxCost int) int {
	cost := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// convert byte to ASCII
		charS := int(s[i]) % 256
		charT := int(t[i]) % 256
		cost[i] = int(math.Abs(float64(charS - charT)))
	}

	maxLen, currentCost := 0, 0
	start := 0
	for end := 0; end < len(s); end++ {
		currentCost = currentCost + cost[end]
		for currentCost > maxCost {
			currentCost = currentCost - cost[start]
			start++
		}

		currentWindow := end - start + 1
		maxLen = max(maxLen, currentWindow)
	}
	return maxLen
}
