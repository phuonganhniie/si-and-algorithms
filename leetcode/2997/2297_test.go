package leetcode_2997

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 3, 4}, 1, 2},
		{[]int{2, 0, 2, 0}, 0, 0},
	}

	for _, tt := range tests {
		got := MinOperations(tt.nums, tt.k)
		if tt.want != got {
			t.Errorf("MinOperations(%v, %v) failed, want %d - got %d", tt.nums, tt.k, tt.want, got)
		}
	}
}
