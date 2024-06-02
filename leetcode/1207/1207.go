package leetcode_1207

func uniqueOccurrences(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	occurMap := make(map[int]int, len(arr))
	for _, num := range arr {
		occurMap[num]++
	}

	countSet := make(map[int]struct{})
	for _, count := range occurMap {
		if _, exist := countSet[count]; exist {
			return false
		}
		countSet[count] = struct{}{}
	}
	return true
}
