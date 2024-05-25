package leetcode_2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	rsList := [2][]int{}
	distinctMap1 := make(map[int]struct{})
	distinctMap2 := make(map[int]struct{})

	for _, num := range nums1 {
		distinctMap1[num] = struct{}{}
	}

	for _, num := range nums2 {
		distinctMap2[num] = struct{}{}
	}

	for num := range distinctMap1 {
		if _, found := distinctMap2[num]; !found {
			rsList[0] = append(rsList[0], num)
		}
	}

	for num := range distinctMap2 {
		if _, found := distinctMap1[num]; !found {
			rsList[1] = append(rsList[1], num)
		}
	}
	return rsList[:]
}
