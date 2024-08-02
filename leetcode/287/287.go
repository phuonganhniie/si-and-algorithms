package leetcode_287

/*
[Medium] 287. Find the Duplicate Number
https://leetcode.com/problems/find-the-duplicate-number/description/?envType=daily-question&envId=2024-03-24
---------------------NOTE---------------------
Time: O(N)
Space: O(1)
Intuition: Fast & slow pointers
*/
func FindDuplicate(nums []int) int {
	var slow, fast = nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

/*
---------------------NOTE---------------------
Time: O(NLogN)
Space: O(N)
Intuition: Binary Search
*/
func FindDuplicateBS(nums []int) int {
	return 0
}
