package leetcode_977

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sqtrArr := make([]int, len(nums))
	start, end, idx := 0, len(nums)-1, len(nums)-1
	for idx >= 0 {
		startNum := nums[start]
		endNum := nums[end]
		if startNum*startNum > endNum*endNum {
			sqtrArr[idx] = startNum * startNum
			start++
		} else {
			sqtrArr[idx] = endNum * endNum
			end--
		}
		idx--
	}
	return sqtrArr
}
