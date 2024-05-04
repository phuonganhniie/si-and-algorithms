package leetcode_167

func TwoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	result := make([]int, 0)

	for numbers[start]+numbers[end] != target {
		if numbers[start]+numbers[end] > target {
			end--
		}
		if numbers[start]+numbers[end] < target {
			start++
		}
	}

	result = append(result, start+1, end+1)
	return result
}
