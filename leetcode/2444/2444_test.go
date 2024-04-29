package leetcode_2444

import "testing"

func TestCountSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		minK int
		maxK int
		want int64
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{[]int{1, 1, 1, 1}, 1, 1, 10},
	}

	for _, tt := range tests {
		got := CountSubarrays(tt.nums, tt.minK, tt.maxK)
		if got != tt.want {
			t.Errorf("CountSubarrays(%v, %v, %v) failed, want %d - got %d", tt.nums, tt.minK, tt.maxK, tt.want, got)
		}
	}
}
