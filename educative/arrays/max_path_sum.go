package arrays

func maxPathSum(arr1 []int, size1 int, arr2 []int, size2 int) int {
	var commonPoints []int
	var ptr1, ptr2 int

	for ptr1 < len(arr1) && ptr2 < len(arr2) {
		if arr1[ptr1] < arr2[ptr2] {
			ptr1++
			continue
		}
		if arr1[ptr1] > arr2[ptr2] {
			ptr2++
			continue
		}
		commonPoints = append(commonPoints, arr1[ptr1])
		ptr1++
		ptr2++
	}

	var maxSum, lastPointArr1, lastPointArr2 int

	for i := 0; i < len(commonPoints); i++ {
		sumArr1 := 0
		for ptr1 := lastPointArr1; i < size1; ptr1++ {
			sumArr1 += arr1[ptr1]
			lastPointArr1 = ptr1
			if arr1[ptr1] == commonPoints[i] {
				lastPointArr1++
				break
			}
		}

		sumArr2 := 0
		for ptr2 := lastPointArr2; i < size2; ptr2++ {
			sumArr2 += arr2[ptr2]
			lastPointArr2 = ptr2
			if arr2[ptr2] == commonPoints[i] {
				lastPointArr2++
				break
			}
		}

		maxPrefixSum := max(sumArr1, sumArr2)
		maxSum += maxPrefixSum
	}

	sumArr1 := 0
	for i := lastPointArr1; i < size1; i++ {
		sumArr1 += arr1[lastPointArr1]
		lastPointArr1++
	}

	sumArr2 := 0
	for i := lastPointArr2; i < size2; i++ {
		sumArr2 += arr2[lastPointArr2]
		lastPointArr2++
	}

	maxPrefixSum := max(sumArr1, sumArr2)
	maxSum += maxPrefixSum
	return maxSum
}

func maxPathSumByEducative(arr1 []int, size1 int, arr2 []int, size2 int) int {
	var sum1, sum2, result int
	var ptr1, ptr2 int
	for ptr1 < size1 && ptr2 < size2 {
		if arr1[ptr1] < arr2[ptr2] {
			sum1 += arr1[ptr1]
			ptr1++
			continue
		}

		if arr1[ptr1] > arr2[ptr2] {
			sum2 += arr2[ptr2]
			ptr2++
			continue
		}

		if arr1[ptr1] == arr2[ptr2] {
			sum1 += arr1[ptr1]
			sum2 += arr2[ptr2]
			result += max(sum1, sum2)
			sum1, sum2 = 0, 0
			ptr1++
			ptr2++
		}
	}

	for ptr1 < size1 {
		sum1 += arr1[ptr1]
		ptr1++
	}

	for ptr2 < size2 {
		sum2 += arr2[ptr2]
		ptr2++
	}

	result += max(sum1, sum2)
	return result
}
