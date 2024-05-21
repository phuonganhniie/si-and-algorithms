package sorting

func SelectionSort(arr []int) []int {
	n := len(arr)
	min := 0

	for i := 0; i < n; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
