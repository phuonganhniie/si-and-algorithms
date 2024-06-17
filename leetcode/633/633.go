package leetcode_633

import "math"

/*
[Medium] 633. Sum of Square Numbers
https://leetcode.com/problems/sum-of-square-numbers/description
Created: 2024-06-17
Done   : mins
Attempt: 1
---------------------NOTE---------------------
Time: O(Sqrt(C))
Space: O(1)
Approach:
*/
func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))
	for a <= b {
		sum := (a * a) + (b * b)
		if sum == c {
			return true
		}
		if sum < c {
			a++
		} else {
			b--
		}
	}
	return false
}
