package two_pointers

func SortedSquares(nums []int) []int {
	n := len(nums)
	sqrtArr := make([]int, n)

	start, end := 0, n-1

	for i := n - 1; i >= 0; i-- {
		sqrt := 0

		if abs(nums[start]) < abs(nums[end]) {
			sqrt = nums[end]
			end--
		} else {
			sqrt = nums[start]
			start++
		}

		sqrtArr[i] = sqrt * sqrt
	}

	return sqrtArr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
