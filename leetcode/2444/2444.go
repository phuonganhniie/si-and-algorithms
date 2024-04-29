package leetcode_2444

func CountSubarrays(nums []int, minK int, maxK int) int64 {
	lastMinK, lastMaxK, validStart := -1, -1, -1
	count := 0

	for i, num := range nums {
		// cập nhật lại validStart khi num < minK || num > maxK
		if num < minK || num > maxK {
			lastMinK, lastMaxK, validStart = -1, -1, i
			continue
		}

		if num == minK {
			lastMinK = i
		}

		if num == maxK {
			lastMaxK = i
		}

		if lastMinK != -1 && lastMaxK != -1 {
			count += min(lastMinK, lastMaxK) - validStart
		}
	}

	return int64(count)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
