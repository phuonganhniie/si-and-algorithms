package leetcode_713

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{1, 2, 3}, 0, 0},
		{[]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19, 18},
	}

	for _, tt := range tests {
		got := NumSubarrayProductLessThanK(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("NumSubarrayProductLessThanK(%v, %v) failed, want %d - got %d", tt.nums, tt.k, tt.want, got)
		}
	}
}
