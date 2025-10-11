package leetcode_167

func TwoSumII(numbers []int, target int) []int {
	result := [2]int{}

	leftPtr, rightPtr := 0, len(numbers)-1
	for leftPtr < rightPtr {
		sum := numbers[leftPtr] + numbers[rightPtr]
		if sum == target {
			result[0] = leftPtr + 1
			result[1] = rightPtr + 1
			break
		}
		if sum <= target {
			leftPtr++
		} else {
			rightPtr--
		}
	}
	return result[:]
}

func TwoSumIIBinarySearch(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		need := target - numbers[i]
		left, right := i+1, n-1

		for left <= right {
			mid := left + (right-left)/2
			if numbers[mid] == need {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid] < need {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return []int{}
}
