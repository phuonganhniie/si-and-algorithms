package leetcode_3025

// Link: https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/?envType=daily-question&envId=2025-09-02

func numberOfPairs(points [][]int) int {
	n := len(points)
	count := 0

	for i := 0; i < n; i++ {
		pointA := points[i]

		for j := 0; j < n; j++ {
			// bỏ qua nếu 2 điểm trùng nhau
			if i == j {
				continue
			}

			pointB := points[j]

			// bỏ qua nếu k thỏa điều kiện: xA <= xB && yA >= yB
			if pointA[0] > pointB[0] || pointA[1] < pointB[1] {
				continue
			}

			valid := true
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}

				pointC := points[k]
				isXContained := pointC[0] >= pointA[0] && pointC[0] <= pointB[0]
				isYContained := pointC[1] <= pointA[1] && pointC[1] >= pointB[1]

				if isXContained && isYContained {
					valid = false
					break
				}
			}

			if valid {
				count++
			}
		}
	}

	return count
}
