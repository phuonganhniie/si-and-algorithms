package leetcode_1482

import "slices"

/*
[Medium] 1482. Minimum Number of Days to Make m Bouquets
https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/description/
Created: 2024-06-20
Done   : 31 mins 30 seconds
Attempt: 3
---------------------NOTE---------------------
Time: O(n * log(max(bloomDay) - min(bloomDay)))
Space: O(1)
Approach: Binary Search
Intuition:
  - Xác định không gian tìm kiếm với giá trị thấp nhất là số ngày bông thứ i nở và số ngày cao nhất trong mảng.
  - Kiểm tra xem trong khoảng thời gian `mid` ngày có thể tạo được bao nhiêu bó hoa được yêu cầu với số bông hoa trên mỗi bó.
  - Ví dụ ở ví dụ 1, từ ngày 1 đến ngày 5 sẽ có 3 bông nở trong khoảng thời gian đó.
    Và vì cần 3 bó mỗi bó 1 bông, nên ta sẽ hạ xuống khoảng tìm kiếm từ ngày 1 đến ngày 3.
    Và trong khoảng thời gian đó vẫn có 3 bông nở sau 3 ngày và đưa về cùng một giá trị, nên kết luận nhỏ nhất là sau 3 ngày ta sẽ bó được 3 bó mỗi bó 1 bông.
*/
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	lowestDay := slices.Min(bloomDay)
	highestDay := slices.Max(bloomDay)
	for lowestDay <= highestDay {
		mid := lowestDay + (highestDay-lowestDay)/2
		if canMakeBouquets(bloomDay, m, k, mid) {
			highestDay = mid - 1
		} else {
			lowestDay = mid + 1
		}
	}
	return lowestDay
}

func canMakeBouquets(bloomDay []int, requiredBouquets, requiredFlowers, days int) bool {
	bouquets, flowers := 0, 0

	for _, day := range bloomDay {
		if day <= days {
			flowers++
			if flowers == requiredFlowers {
				bouquets++
				flowers = 0
			}
		}
		if day > days {
			flowers = 0
		}
	}
	return bouquets >= requiredBouquets
}
