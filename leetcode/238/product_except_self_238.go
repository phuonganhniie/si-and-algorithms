package leetcode

import "fmt"

func ProductExceptSelf(nums []int) []int {
	n := len(nums)

	answer := make([]int, n)
	leftProd := make([]int, n)
	rightProd := make([]int, n)

	leftProd[0] = 1
	for i := 1; i < n; i++ {
		leftProd[i] = leftProd[i-1] * nums[i-1]
		fmt.Println("left prod[i]: ", leftProd[i])
	}

	rightProd[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rightProd[i] = rightProd[i+1] * nums[i+1]
		fmt.Println("right prod[i]: ", rightProd[i])
	}

	for i := 0; i < n; i++ {
		answer[i] = leftProd[i] * rightProd[i]
	}
	return answer
}
