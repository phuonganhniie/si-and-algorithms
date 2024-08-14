package sorting

func partition01(arr []int, size int) ([]int, int) {
	start, end := 0, size-1
	countSwap := 0

	for start < end {
		if arr[start] == 0 && arr[end] == 1 {
			start++
			end--
			continue
		}

		if arr[start] == 1 && arr[end] == 1 {
			end--
			continue
		}

		if arr[start] == 0 && arr[end] == 0 {
			start++
			continue
		}

		if arr[start] == 1 && arr[end] == 0 {
			arr[start], arr[end] = arr[end], arr[start]
			countSwap++
			end--
			continue
		}
	}
	return arr, countSwap
}

func partition01Optimize(arr []int, size int) ([]int, int) {
	start, end := 0, size-1
	countSwap := 0

	for start < end {
		for arr[start] == 0 {
			start++
		}

		for arr[end] == 1 {
			end--
		}

		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
			countSwap++
		}
	}
	return arr, countSwap
}
