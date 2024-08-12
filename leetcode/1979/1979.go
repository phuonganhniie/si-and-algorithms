package leetcode

/*
[Easy] 1979. Find Greatest Common Divisor of Array
https://leetcode.com/problems/find-greatest-common-divisor-of-array/description/
Created: 2024-08-12
Done   : 04 mins 58 seconds
Attempt: 2
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: Euclid's math algorithm
*/
func findGCD(nums []int) int {
	s, l := nums[0], nums[0]

	for _, num := range nums {
		s = min(s, num)
	}

	for _, num := range nums {
		l = max(l, num)
	}

	rs := GCD(nums, s, l)
	return rs
}

func GCD(nums []int, s, l int) int {
	for l != 0 {
		s, l = l, s%l
	}
	return s
}
