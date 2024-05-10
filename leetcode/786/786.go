package leetcode_786

func KthSmallestPrimeFraction(arr []int, k int) []int {
	low, high := 0.0, 1.0
	n := len(arr)
	bestNumerator, bestDenominator := 0, 0

	for low < high {
		mid := low + (high-low)/2
		count, maxNumerator, denominator := 0, 0, 1
		den := 1

		for num := 0; num < n; num++ {
			for den < n && float64(arr[num])/float64(arr[den]) > mid {
				den++
			}

			// if a valid fraction is found, count it
			if den < n {
				count += (n - den)
				if maxNumerator*arr[den] < denominator*arr[num] {
					maxNumerator = arr[num]
					denominator = arr[den]
				}
			}
		}

		if count == k {
			bestNumerator, bestDenominator = maxNumerator, denominator
			break
		}
		if count < k {
			low = mid
		} else {
			high = mid
		}
	}

	return []int{bestNumerator, bestDenominator}
}
