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
