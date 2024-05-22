package leetcode_283

func moveZeroes(nums []int) []int {
	if len(nums) == 1 && nums[0] == 0 {
		return nums
	}

	low := 0
	for high := 0; high < len(nums); high++ {
		if nums[high] != 0 {
			nums[high], nums[low] = nums[low], nums[high]
			low++
		}
	}
	return nums
}
