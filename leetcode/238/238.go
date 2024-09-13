package leetcode_238

func ProductExceptSelf(nums []int) []int {
	totalLeftProd := make([]int, len(nums))
	totalRightProd := make([]int, len(nums))

	totalLeftProd[0] = 1
	for i := 1; i < len(nums); i++ {
		totalLeftProd[i] = totalLeftProd[i-1] * nums[i-1]
	}

	totalRightProd[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		totalRightProd[i] = totalRightProd[i+1] * nums[i+1]
	}

	rs := []int{}
	ptr1, ptr2 := 0, 0
	for ptr1 < len(totalLeftProd) && ptr2 < len(totalRightProd) {
		rs = append(rs, totalLeftProd[ptr1]*totalRightProd[ptr2])
		ptr1++
		ptr2++
	}

	return rs
}
