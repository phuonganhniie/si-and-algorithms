package codility

import "math"

func SolutionLesson3_3(A []int) int {
	if len(A) < 2 {
		return 0
	}

	sumFront, sumBack := 0, sum(A)
	diffList := []int{}

	for _, num := range A {
		sumFront += num
		sumBack -= num

		absSum := int(math.Abs(float64(sumFront) - float64(sumBack)))
		diffList = append(diffList, absSum)
	}

	rs := findMin(diffList)
	return rs
}

func findMin(arr []int) int {
	minNum := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < minNum {
			minNum = arr[i]
		}
	}
	return minNum
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// =======================================

func SolutionLesson3_3a(A []int) int {
	if len(A) < 2 {
		return 0
	}

	sumFront, sumBack := A[0], 0
	for i := 1; i < len(A); i++ {
		sumBack += A[i]
	}

	minDiff := int(math.Abs(float64(sumFront) - float64(sumBack)))

	for i := 1; i < len(A)-1; i++ { // Iterate until the second to last element
		sumFront += A[i]
		sumBack -= A[i]
		diff := int(math.Abs(float64(sumFront) - float64(sumBack)))
		if diff == 0 {
			return 0 // Early exit for minimal difference
		}

		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
