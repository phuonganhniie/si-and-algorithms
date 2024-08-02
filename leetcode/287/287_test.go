package leetcode_287

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}, 9},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
	}

	for _, tt := range tests {
		got := FindDuplicate(tt.nums)
		if got != tt.want {
			t.Errorf("FindDuplicate(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestFindDuplicateBS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
	}

	for _, tt := range tests {
		got := FindDuplicateBS(tt.nums)
		if got != tt.want {
			t.Errorf("FindDuplicate(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
