package arrays

func smallestPositiveMissingNumber(arr []int, size int) int {
	trackMap := make(map[int]bool)
	for _, num := range arr {
		trackMap[num] = true
	}

	missingNum := -1
	for i := 1; i <= size; i++ {
		if !trackMap[i] {
			missingNum = i
		}
	}
	return missingNum
}
