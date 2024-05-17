package arrays

func BinarySearch(data []int, value int) bool {
	low, high := 0, len(data)-1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] == value {
			return true
		}

		if data[mid] <= value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
