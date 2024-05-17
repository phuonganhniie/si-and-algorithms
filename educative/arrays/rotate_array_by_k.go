package arrays

func RotateArray(data []int, k int) []int {
	tempArr := make([]int, k)
	for i := 0; i < k; i++ {
		tempArr[i] = data[i]
	}

	rotatedArr := data[k:]
	rotatedArr = append(rotatedArr, tempArr...)

	return rotatedArr
}
