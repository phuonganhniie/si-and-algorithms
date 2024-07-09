package leetcode_1823

/*
[Medium] 1518. Water Bottles
https://leetcode.com/problems/water-bottles/description
Created: 2024-07-09
Done   : 11 mins 03s
Attempt: 1
---------------------NOTE---------------------
Time: O(1)
Space: O(1)
Approach: Math - Josephus Problem
*/
func findTheWinner(n int, k int) int {
	// res := 0
	// for i := 2; i <= n; i++ {
	//     res = (res + k) % i
	// }
	// return res + 1

	if n > 1 {
		return (findTheWinner(n-1, k)+k-1)%n + 1
	}
	return 1
}
