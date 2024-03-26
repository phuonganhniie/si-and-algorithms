package leetcode_41

import "testing"

func TestFirstMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for _, tt := range tests {
		got := FirstMissingNumber2(tt.nums)
		if got != tt.want {
			t.Errorf("FirstMissingNumber(%v) failed, want %d - got %d", tt.nums, tt.want, got)
		}
	}
}
