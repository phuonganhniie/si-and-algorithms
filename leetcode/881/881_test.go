package leetcode_881

import "testing"

func TestNumRescueBoats(t *testing.T) {
	cases := []struct {
		people []int
		limit  int
		want   int
	}{
		{[]int{2, 2}, 6, 1},
		{[]int{3, 2, 3, 2, 2}, 6, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
		{[]int{3, 1, 7}, 7, 2},
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 2, 1}, 3, 3},
	}

	for _, c := range cases {
		got := numRescueBoats(c.people, c.limit)
		if got != c.want {
			t.Errorf("NumRescueBoats(%v, %v) = %d, want %d got %d", c.people, c.limit, got, c.want, got)
		}
	}
}
