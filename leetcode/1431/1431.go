package leetcode_1431

func KidsWithCandies(candies []int, extraCandies int) []bool {
	rs := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		largestTemp := candies[i] + extraCandies
		for j := 0; j < len(candies); j++ {
			if largestTemp < candies[j] {
				rs[i] = false
				break
			}
			rs[i] = true
		}
	}
	return rs
}

func KidsWithCandiesOptimized(candies []int, extraCandies int) []bool {
	rs := make([]bool, len(candies))

	// find the maximum candies that any kid is hold
	maxCandies := 0
	for _, candy := range candies {
		if candy >= maxCandies {
			maxCandies = candy
		}
	}

	for i, candy := range candies {
		largestTemp := candy + extraCandies
		if largestTemp < maxCandies {
			rs[i] = false
			continue
		}
		rs[i] = true
	}
	return rs
}
