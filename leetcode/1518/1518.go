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
func numWaterBottles2(soChaiDangCo int, soVoCanDoi1Chai int) int {
	daUong := 0

	for soChaiDangCo >= soVoCanDoi1Chai {
		soChaiDangCo -= soVoCanDoi1Chai
		daUong += soVoCanDoi1Chai

		// Doi vo lay 1 chai moi
		soChaiDangCo++
	}

	return daUong + soChaiDangCo
}

/*
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: Cách của empe, ún xong một lần rồi chạy đi đổi
*/
func numWaterBottles3(numBottles int, numExchange int) int {
	totalDrank := 0
	full := numBottles
	empties := 0

	for full > 0 {
		// Uống hết số chai đầy hiện có
		totalDrank += full
		empties += full

		// Đổi vỏ lấy chai đầy mới
		full = empties / numExchange
		empties = empties % numExchange
	}

	return totalDrank
}
