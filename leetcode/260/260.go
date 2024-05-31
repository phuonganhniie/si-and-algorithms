package leetcode_260

// Solution 1: Brute-force
// Time: O(n) - Space: O(n)
func singleNumber(nums []int) []int {
	trackingMap := make(map[int]int)
	for _, num := range nums {
		trackingMap[num]++
	}

	result := []int{}
	for num, count := range trackingMap {
		if count == 1 {
			result = append(result, num)
		}
	}
	return result
}

// Solution 2: Bit manipulation
// Time: O(n) - Space: O(1)
func singleNumberOptimize(nums []int) []int {
	all := 0
	for _, num := range nums {
		all ^= num
	}

	setBit := all & -all
	a, b := 0, 0
	for _, num := range nums {
		if num&setBit == 0 {
			b ^= num
		} else {
			a ^= num
		}
	}
	return []int{a, b}
}
