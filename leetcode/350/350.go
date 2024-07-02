package leetcode_350

import "sort"

/*
[Medium] 350. Intersection of Two Arrays II
https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
Created: 2024-07-02
Done   : 13 mins
Attempt: 1
---------------------NOTE---------------------
Time: O(nlogn)
Space: O(min(m,n))
Approach: 2 pointers + sorting
*/
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	ans := make([]int, 0)
	ptr1, ptr2 := 0, 0
	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] == nums2[ptr2] {
			ans = append(ans, nums1[ptr1])
			ptr1++
			ptr2++
			continue
		}
		if nums1[ptr1] < nums2[ptr2] {
			ptr1++
			continue
		}
		ptr2++
	}
	return ans
}
