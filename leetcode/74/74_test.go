package leetcode_74

import "testing"

func TestSearchMatrix(t *testing.T) {
	cases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}

	for _, c := range cases {
		got := SearchMatrix(c.matrix, c.target)
		if got != c.want {
			t.Errorf("SearchMatrix test failed, want %v - got %v", c.want, got)
		}
	}
}
