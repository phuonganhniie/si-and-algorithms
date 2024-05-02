package leetcode_2441

import "testing"

func TestFindMaxK(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{-1, 2, -3, 3}, 3},
		{[]int{-1, 10, 6, 7, -7, 1}, 7},
		{[]int{-10, 8, 6, 7, -2, -3}, -1},
	}

	for _, c := range cases {
		got := FindMaxK(c.nums)
		if got != c.want {
			t.Errorf("FindMaxK(%v) = %v, want %v got %v", c.nums, got, c.want, got)
		}
	}
}
