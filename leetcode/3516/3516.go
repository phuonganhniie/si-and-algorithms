package leetcode_3516

import "math"

// Link: https://leetcode.com/problems/find-closest-person/?envType=daily-question&envId=2025-09-04

func findClosest(x int, y int, z int) int {
	dxz := int(math.Abs(float64(x-z)))
	dyz := int(math.Abs(float64(y-z)))

	if dxz < dyz {
		return 1
	} 
	
	if dxz > dyz {
		return 2
	} 
		
	return 0
}