package sorting

func SelectionSort(arr []int) []int {
	n := len(arr)
	min, temp := 0, 0

	for i := 0; i < n; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}

	return arr
}
