package leetcode

/*
[Medium] 01. Two Sum
https://leetcode.com/problems/two-sum/description/
Created: 2024-07-13
Done   :
Attempt: 3
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: HashMap
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if mapNumIdx, exist := m[complement]; exist {
			return []int{mapNumIdx, i}
		}
		m[num] = i
	}
	return []int{}
}
