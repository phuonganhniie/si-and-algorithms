package leetcode_1304

// Link: https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/?envType=daily-question&envId=2025-09-07

func sumZero(n int) []int {
	result := make([]int, 0, n)

	for i := 1; i <= n/2; i++ {
		result = append(result, i, -i)
	}

	if n%2 == 1 {
		result = append(result, 0)
	}

	return result
}

func sumZero2(n int) []int{
	result := make([]int, 0, n)

	// chọn n-1 số liên tiếp từ 1...(n-1)
	sum := 0
	for num := 1; num < n; num++ {
		result = append(result, num)
		sum += num
	}

	result = append(result, -sum)
	return result
}