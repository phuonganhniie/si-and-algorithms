package leetcode_35

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, c := range cases {
		got := SearchInsert(c.nums, c.target)
		if got != c.want {
			t.Errorf("SearchInsert test failed, want %d - got %d", c.want, got)
		}
	}
}
