package two_pointers

import "fmt"

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	trackingWindow := 0
	product := 1
	left := 0

	for right, num := range nums {
		product *= num

		// If product is equal or larger than k, move the left pointer to make the product smaller
		for product >= k {
			product /= nums[left]
			left++
		}
		fmt.Printf("Window [%d:%d] - ", left, right)
		fmt.Printf("Elements: %v, ", nums[left:right+1])
		fmt.Printf("Product: %d\n", product)
		trackingWindow += right - left + 1
	}
	return trackingWindow
}
