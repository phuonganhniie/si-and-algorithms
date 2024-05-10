package leetcode_786

import (
	"reflect"
	"testing"
)

func TestKthSmallestPrimeFraction(t *testing.T) {
	cases := []struct {
		arr  []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 5}, 3, []int{2, 5}},
		{[]int{1, 7}, 1, []int{1, 7}},
	}

	for _, c := range cases {
		got := KthSmallestPrimeFraction(c.arr, c.k)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("KthSmallestPrimeFraction test failed, want %v - got %v", c.want, got)
		}
	}
}
