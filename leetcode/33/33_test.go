package leetcode_33

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 1, 5},
	}

	for _, tt := range cases {
		got := Search(tt.nums, tt.target)
		if tt.want != got {
			t.Errorf("Search(%v, %v) failed, want %d - got %d", tt.nums, tt.target, tt.want, got)
		}
	}
}
