package leetcode_704

import "testing"

func TestSearchBasic(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, c := range cases {
		got := BasicSearch(c.nums, c.target)
		if got != c.want {
			t.Errorf("BasicSearch(%v, %v) = %v, want %d, got %d", c.nums, c.target, got, c.want, got)
		}
	}
}
