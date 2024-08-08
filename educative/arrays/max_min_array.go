package arrays

func maxMinArr(arr []int, size int) []int {
	start, end := 0, size-1
	tempArr := make([]int, size)
	copy(tempArr, arr)

	for i := 0; i < size; i++ {
		if i%2 != 0 {
			arr[i] = tempArr[start]
			start++
		} else {
			arr[i] = tempArr[end]
			end--
		}
	}
	return arr
}
