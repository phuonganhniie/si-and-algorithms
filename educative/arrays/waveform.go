package arrays

func swap(data []int, x, y int) {
	data[x], data[y] = data[y], data[x]
}

func WaveArray(arr []int) []int {
	for i := 1; i < len(arr); i += 2 {
		if i-1 >= 0 && arr[i] > arr[i-1] {
			swap(arr, i, i-1)
		}
		if i+1 < len(arr) && arr[i] > arr[i+1] {
			swap(arr, i, i+1)
		}
	}
	return arr
}
