package binarysearchmodified

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 8, 4, 5, 6, -1, 9}, 5},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{3, 2}, 1},
	}

	countErr := 0

	for _, tt := range tests {
		got := longestSubsequenceBS(tt.nums)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("longestSubsequence test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}

func BenchmarkLongestSubsequenceBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestSubsequenceBS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	}
}
