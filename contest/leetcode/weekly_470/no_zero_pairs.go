package weekly_470

import (
	"strconv"
)

// CountNoZeroPairs đếm số cặp (a, b) sao cho a + b = n và cả a, b đều không chứa chữ số 0
// Sử dụng Digit DP để tối ưu hóa
func CountNoZeroPairs(n int64) int64 {
	// Tạo biến để lưu input
	trivanople := n

	// Chuyển n thành mảng các chữ số (đảo ngược để dễ xử lý từ phải sang trái)
	nStr := strconv.FormatInt(trivanople, 10)
	lenN := len(nStr)
	digits := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		digits[lenN-1-i] = int(nStr[i] - '0')
	}

	// Memoization cache
	cache := make(map[string]int64)

	// Hàm đệ quy với memoization
	var solve func(pos, carry, lenA, lenB int) int64
	solve = func(pos, carry, lenA, lenB int) int64 {
		if pos == lenN {
			if carry == 0 {
				return 1
			}
			return 0
		}

		// Tạo key cho cache
		key := strconv.Itoa(pos) + "," + strconv.Itoa(carry) + "," + strconv.Itoa(lenA) + "," + strconv.Itoa(lenB)
		if val, exists := cache[key]; exists {
			return val
		}

		// Xác định range cho chữ số của a và b
		var rangeA, rangeB []int
		if pos < lenA {
			rangeA = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		} else {
			rangeA = []int{0}
		}

		if pos < lenB {
			rangeB = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		} else {
			rangeB = []int{0}
		}

		numWays := int64(0)
		for _, da := range rangeA {
			for _, db := range rangeB {
				summ := da + db + carry
				if summ%10 == digits[pos] {
					numWays += solve(pos+1, summ/10, lenA, lenB)
				}
			}
		}

		cache[key] = numWays
		return numWays
	}

	totalPairs := int64(0)
	for lenA := 1; lenA <= lenN; lenA++ {
		for lenB := 1; lenB <= lenN; lenB++ {
			totalPairs += solve(0, 0, lenA, lenB)
		}
	}

	return totalPairs
}

// hasZero kiểm tra xem một số có chứa chữ số 0 hay không
func hasZero(num int64) bool {
	if num == 0 {
		return true
	}
	for num > 0 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}

// generateNoZeroNumbers sinh ra tất cả các số không chứa chữ số 0 từ 1 đến maxNum
func generateNoZeroNumbers(maxNum int64) []int64 {
	result := []int64{}

	// Sử dụng BFS để sinh số
	queue := []int64{}
	for digit := int64(1); digit <= 9; digit++ {
		if digit <= maxNum {
			queue = append(queue, digit)
			result = append(result, digit)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for digit := int64(1); digit <= 9; digit++ {
			next := current*10 + digit
			if next <= maxNum {
				queue = append(queue, next)
				result = append(result, next)
			}
		}
	}

	return result
}

// CountNoZeroPairsOld là phương pháp cũ (không tối ưu cho n lớn)
// Giữ lại để tham khảo và test với số nhỏ
func CountNoZeroPairsOld(n int64) int64 {
	noZeroNums := generateNoZeroNumbers(n - 1)
	count := int64(0)
	for _, a := range noZeroNums {
		b := n - a
		if b > 0 && !hasZero(b) {
			count++
		}
	}
	return count
}
