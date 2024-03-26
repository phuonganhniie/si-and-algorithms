package leetcode_442

import "sort"

func FindDuplicates(nums []int) []int {
	result := make([]int, 0)

	n := len(nums)
	trackingMap := make(map[int]int, n)

	for _, num := range nums {
		trackingMap[num]++
	}

	for num, count := range trackingMap {
		if count >= 2 {
			result = append(result, num)
		}
	}

	sort.Ints(result)

	return result
}

/**
 * Chôm solution
 * Để ý là giá trị từ 1 - n và 1 phần tử chỉ có tối đa 1 lần lặp
 * => dùng trick đổi dấu phần tử theo chỉ số nếu được sort
 */
func FindDuplicates2(nums []int) []int {
	var result []int

	for _, num := range nums {
		idx := abs(num) - 1

		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			result = append(result, abs(num))
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
