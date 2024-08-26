package twopointers

func productExceptSelfNaive(arr []int) []int {
	result := make([]int, len(arr), len(arr))

	for i := 0; i < len(arr); i++ {
		leftPrd := arr[:i]
		rightPrd := arr[i+1:]

		leftPrex := 1
		for _, num := range leftPrd {
			leftPrex *= num
		}

		rightPrex := 1
		for _, num := range rightPrd {
			rightPrex *= num
		}

		result[i] = leftPrex * rightPrex
	}
	return result
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	rs := make([]int, n)

	for i := range rs {
		rs[i] = 1
	}

	leftPrefix := 1
	for i := 0; i < n; i++ {
		rs[i] *= leftPrefix
		leftPrefix *= nums[i]
	}

	rightSuffix := 1
	for i := n - 1; i >= 0; i-- {
		rs[i] *= rightSuffix
		rightSuffix *= nums[i]
	}

	return rs
}

func productExceptSelfEducative(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := range res {
		res[i] = 1
	}

	leftPrefix, rightSuffix := 1, 1
	left, right := 0, n-1

	for left < n && right > -1 {
		res[left] *= leftPrefix
		res[right] *= rightSuffix

		leftPrefix *= nums[left]
		rightSuffix *= nums[right]

		left++
		right--
	}

	return res
}
