package leetcode_162

import "testing"

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{5, 4, 3, 4, 5}, 0},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{2, 1}, 0},
		{[]int{1}, 0},
		{[]int{1, 2, 3, 1}, 2},
	}

	for _, c := range cases {
		got := findPeakElement(c.nums)
		if got != c.want {
			t.Errorf("FindPeakElement test failed, want %d - got %d", c.want, got)
		}
	}
}
