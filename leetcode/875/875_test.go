package leetcode_875

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	cases := []struct {
		piles     []int
		limitHour int
		want      int
	}{
		{[]int{3, 6, 7, 100, 1, 1, 1, 1, 1}, 10, 50},
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}

	for _, c := range cases {
		got := MinEatingSpeed(c.piles, c.limitHour)
		if got != c.want {
			t.Errorf("MinEatingSpeed test failed, want %d - got %d", c.want, got)
		}
	}
}
