package leetcode_167

func TwoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	result := make([]int, 0)

	for start < end {
		totalSum := numbers[start] + numbers[end]
		if totalSum == target {
			result = append(result, start+1, end+1)
		}
		if totalSum < target {
			start++
		} else {
			end--
		}
	}
	return result
}
