package leetcode_1550

func threeConsecutiveOdds(arr []int) bool {
	countOdds := 0
	for _, num := range arr {
		if num%2 != 0 {
			countOdds++
		} else {
			countOdds = 0
		}

		if countOdds == 3 {
			return true
		}
	}
	return false
}
