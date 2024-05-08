package leetcode_1011

import "testing"

func TestShipWithinDays(t *testing.T) {
	cases := []struct {
		weights []int
		days    int
		want    int
	}{
		{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
		{[]int{1, 2, 3, 1, 1}, 4, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
	}

	for _, c := range cases {
		got := ShipWithinDays(c.weights, c.days)
		if got != c.want {
			t.Errorf("ShipWithinDays test failed, want %d - got %d", c.want, got)
		}
	}
}
