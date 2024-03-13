package leetcode_2485

/**
 * My Solution - Binary Search approach
 * Complexity: O(log n)
 */
func binarySearchPivot(n int, totalSum int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)/2
		leftSum := mid * (mid + 1) / 2
		rightSum := totalSum - leftSum + mid

		if leftSum == rightSum {
			return mid
		} else if leftSum < rightSum {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func PivotIntegerBS(n int) int {
	totalSum := n * (n + 1) / 2
	return binarySearchPivot(n, totalSum)
}

/**
 * ChatGPT Solution - Prefix Sum approach
 * Complexity: O(n)
 */
func PivotInteger(n int) int {
	total := n * (n + 1) / 2
	l := 0

	for i := 1; i <= n; i++ {
		r := total - l - i

		if l == r {
			return i
		}

		l += i
	}

	return -1
}

/**
 * DucChu Solution - Công thức cấp số cộng + phương trình
 * x = sqrt((N ^ 2 + N) / 2)
 */
