package leetcode

func TwoSum(nums []int, target int) []int {
	compareMap := make(map[int]int, 10)

	for i, num := range nums {
		complement := target - num

		if mapIndx, found := compareMap[complement]; found {
			return []int{mapIndx, i}
		}
		compareMap[num] = i
	}
	return []int{}
}
