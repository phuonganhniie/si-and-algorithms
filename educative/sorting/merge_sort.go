package sorting

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	start, end := 0, len(arr)-1
	mid := start + (end-start)/2
	leftArr := arr[start : mid+1]
	rightArr := arr[mid+1 : end+1]

	leftArr = MergeSort(leftArr)
	rightArr = MergeSort(rightArr)

	return merge(leftArr, rightArr)
}

func merge(leftArr []int, rightArr []int) []int {
	merged := make([]int, 0, len(leftArr)+len(rightArr))
	l, r := 0, 0

	for l < len(leftArr) && r < len(rightArr) {
		if leftArr[l] < rightArr[r] {
			merged = append(merged, leftArr[l])
			l++
		} else {
			merged = append(merged, rightArr[r])
			r++
		}
	}

	merged = append(merged, leftArr[l:]...)
	merged = append(merged, rightArr[r:]...)

	return merged
}
