package arrays

func arrayIndexMaxDiff(arr []int, size int) int {
	maxDiff := 0
	for first := 0; first < size; first++ {
		for sec := 0; sec < size; sec++ {
			if arr[first] < arr[sec] {
				maxDis := sec - first
				maxDiff = max(maxDiff, maxDis)
			}
		}
	}
	return maxDiff
}
