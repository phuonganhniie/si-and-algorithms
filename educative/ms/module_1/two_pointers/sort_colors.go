package twopointers

func sortColors(colors []int) []int {
	const maxNum = 2
	countArr := make([]int, maxNum+1, maxNum+1)
	for _, num := range colors {
		countArr[num]++
	}

	sorted := []int{}
	for i := 0; i <= maxNum; i++ {
		for range countArr[i] {
			sorted = append(sorted, i)
		}
	}
	return sorted
}
