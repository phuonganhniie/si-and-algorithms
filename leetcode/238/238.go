package leetcode_238

func ProductExceptSelf(nums []int) []int {
	leftProductArr := make([]int, len(nums))
	rightProductArr := make([]int, len(nums))
	result := make([]int, len(nums))

	leftProductArr[0] = 1
	rightProductArr[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		leftProductArr[i] = leftProductArr[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		rightProductArr[i] = rightProductArr[i+1] * nums[i+1]
	}

	leftPtr, rightPtr := 0, 0
	for leftPtr < len(leftProductArr) && rightPtr < len(rightProductArr) {
		numInLeft := leftProductArr[leftPtr]
		numInRight := rightProductArr[rightPtr]
		result[leftPtr] = numInLeft * numInRight

		leftPtr++
		rightPtr++
	}
	return result
}
