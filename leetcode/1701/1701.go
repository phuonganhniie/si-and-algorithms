package leetcode_1701

/*
[Medium] 1701. Average Waiting Time
https://leetcode.com/problems/average-waiting-time/description/
Created: 2024-07-09
Done   : 08 mins 43s
Attempt: 3
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: Simulator -> suy luận
*/
func averageWaitingTime(customers [][]int) float64 {
	idleTime, waitTime := 0, 0
	for i := 0; i < len(customers); i++ {
		idleTime = max(customers[i][0], idleTime) + customers[i][1]
		waitTime += idleTime - customers[i][0]
	}

	average := float64(waitTime) / float64(len(customers))
	return average
}

func averageWaitingTime2(customers [][]int) float64 {
	totalWaitingTime := 0
	chefAvailableTime := customers[0][0] // Thời điểm mà đầu bếp sẽ rảnh sau khi hoàn thành đơn hàng hiện tại

	for _, customer := range customers {
		arrivalTime := customer[0] // Thời gian đến của khách hàng
		cookingTime := customer[1] // Thời gian cần thiết để nấu đơn hàng của khách hàng

		// Tính toán thời gian chờ của khách hàng
		waitTime := chefAvailableTime - arrivalTime

		// Nếu đầu bếp rảnh trước khi khách hàng đến, khách hàng không phải đợi
		if waitTime < 0 {
			waitTime = 0
			chefAvailableTime = arrivalTime
		}

		// Cập nhật thời gian đầu bếp sẽ rảnh sau khi hoàn thành đơn hàng của khách hàng hiện tại
		chefAvailableTime += cookingTime

		// Cộng thời gian chờ của khách hàng hiện tại vào tổng thời gian chờ
		totalWaitingTime += waitTime + cookingTime
	}

	return float64(totalWaitingTime) / float64(len(customers))
}
