package leetcode_3075

import (
	"reflect"
	"testing"
)

func TestMaximumHappinessSum(t *testing.T) {
	cases := []struct {
		happiness []int
		k         int
		want      int64
	}{
		{[]int{1, 2, 3}, 2, 4},
		{[]int{1, 1, 1, 1}, 2, 1},
		{[]int{2, 3, 4, 5}, 1, 5},
	}

	for _, c := range cases {
		got := MaximumHappinessSum(c.happiness, c.k)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("RemoveNodes failed, want %v - got %v", c.want, got)
		}
	}
}

func TestMaximumHappinessSum2(t *testing.T) {
	cases := []struct {
		happiness []int
		k         int
		want      int64
	}{
		{[]int{0, 0, 0}, 3, 0},
		{[]int{1, 2, 3}, 2, 4},
		{[]int{1, 1, 1, 1}, 2, 1},
		{[]int{2, 3, 4, 5}, 1, 5},
		{[]int{0}, 1, 0},
	}

	for _, c := range cases {
		got := MaximumHappinessSum2(c.happiness, c.k)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("RemoveNodes failed, want %v - got %v", c.want, got)
		}
	}
}
