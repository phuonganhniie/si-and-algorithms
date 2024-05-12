package leetcode_167

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 18, []int{2, 3}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{5, 25, 75}, 100, []int{2, 3}},
		{[]int{2, 7, 11, 13, 15}, 15, []int{1, 4}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, c := range cases {
		got := TwoSum(c.numbers, c.target)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("TwoSum test fail, want %v - got %v", c.want, got)
		}
	}
}
