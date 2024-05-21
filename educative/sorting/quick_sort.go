package sorting

func QuickSort(arr []int) []int {
	start, end := 0, len(arr)-1
	return quickSort(arr, start, end)
}

func quickSort(arr []int, start, end int) []int {
	if end <= start {
		return arr
	}

	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)

	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // choosing the last element as pivot
	l := low - 1       // start from index -1 OR index of smaller element

	for r := low; r < high; r++ {
		if arr[r] < pivot {
			l++
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	// swap the value at index l+1 with the value of pivot
	arr[l+1], arr[high] = arr[high], arr[l+1]

	return l + 1
}
