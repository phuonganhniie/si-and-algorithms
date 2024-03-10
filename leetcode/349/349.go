package leetcode_349

func Intersection(nums1 []int, nums2 []int) []int {
	intersectMap := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		if _, ok := intersectMap[num]; !ok {
			intersectMap[num] = true
		}
	}

	for _, num := range nums2 {
		if v, ok := intersectMap[num]; ok && v {
			result = append(result, num)
			intersectMap[num] = false
		}
	}

	return result
}
