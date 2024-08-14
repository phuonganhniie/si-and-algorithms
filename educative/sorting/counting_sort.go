package sorting

func countingSort(arr []int) []int {
	var maxNum int
	for _, num := range arr {
		maxNum = max(maxNum, num)
	}

	countArr := make([]int, maxNum+1, maxNum+1)
	for _, num := range arr {
		countArr[num]++
	}

	rs := []int{}
	for i := 0; i <= maxNum; i++ {
		if countArr[i] != 0 {
			for range countArr[i] {
				rs = append(rs, i)
			}
		}
	}

	copy(arr, rs)
	return arr
}
