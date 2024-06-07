package leetcode_1442

import (
	"fmt"
)

func countTriplets(arr []int) int {
	n := len(arr)
	count := 0

	// tạo một mảng prefix XOR để lưu trữ tổng XOR cho mỗi vị trí, các phần tử ban đầu đều bằng 0
	prefixXOR := make([]int, n+1)

	// tính prefix XOR cho mỗi phần tử trong mảng
	for i := 0; i < n; i++ {
		prefixXOR[i+1] = prefixXOR[i] ^ arr[i]
	}

	// duyệt từng cặp (i,k) để kiểm tra điều kiện tổng i^k = 0
	result := [][]int{}
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if prefixXOR[i] == prefixXOR[k+1] {
				count += (k - i)
				result = append(result, []int{i, k - i, k})
			}
		}
	}
	fmt.Printf("Triplets are\n: %v", result)
	return count
}
