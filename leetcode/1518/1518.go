package leetcode_1518

/*
[Medium] 1518. Water Bottles
https://leetcode.com/problems/water-bottles/description
Created: 2024-07-07
Done   : 11 mins 03s
Attempt: 1
---------------------NOTE---------------------
Time: O(1)
Space: O(1)
Approach: Math
*/
func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}

/*
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: Simulation (này tự nghĩ ý tưởng nên cũng k biết theo pattern gì)
*/
func numWaterBottles2(numBottles int, numExchange int) int {
	consumedBottles := 0

	for numBottles >= numExchange {
		numBottles -= numExchange
		consumedBottles += numExchange

		// Exchange them for one full bottle
		numBottles++
	}

	return consumedBottles + numBottles
}
