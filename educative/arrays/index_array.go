package arrays

func IndexArray(arr []int, size int) []int {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != -1 && arr[i] != i {
			temp = arr[i]
			arr[i] = arr[temp]
			arr[temp] = temp
		}
	}
	return arr
}

func IndexArray2(arr []int, size int) []int {
	existNums := make(map[int]bool)
	for _, num := range arr {
		if num != -1 {
			existNums[num] = true
		}
	}

	rs := make([]int, size)
	for i := 0; i < size; i++ {
		if existNums[i] {
			rs[i] = i
			continue
		}
		rs[i] = -1
	}
	copy(arr, rs)
	return arr
}
