package leetcode_2037

import (
	"sort"
)

/*
[Easy] 2037. Minimum Number of Moves to Seat Everyone
https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description/
Created: 2024-06-13
Done   : 30 mins 52 seconds
Attempt: 1
---------------------NOTE---------------------
Time: O(n log n)
Space: O(1)
Approach: Này cũng chưa biết là dạng gì, làm theo bản năng thôi
*/
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	totalMoves := 0
	studentPtr, seatPtr := 0, 0
	for studentPtr < len(students) && seatPtr < len(seats) {
		seat := seats[seatPtr]
		student := students[studentPtr]
		if student > seat {
			totalMoves += student - seat
		} else {
			totalMoves += seat - student
		}
		studentPtr++
		seatPtr++
	}
	return totalMoves
}
