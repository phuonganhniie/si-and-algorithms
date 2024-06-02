package leetcode_334

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	firstNum, secondNum := 1<<63-1, 1<<63-1
	for _, num := range nums {
		if num <= firstNum {
			firstNum = num
			continue
		}
		if num <= secondNum {
			secondNum = num
			continue
		}
		return true
	}
	return false
}
