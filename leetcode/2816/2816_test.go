package leetcode_2816

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestDoubleIt(t *testing.T) {
	cases := []struct {
		head []int
		want []int
	}{
		{[]int{1, 8, 1}, []int{3, 6, 2}},
		{[]int{1, 8, 9}, []int{3, 7, 8}},
		{[]int{9, 9, 9}, []int{1, 9, 9, 8}},
	}

	for _, c := range cases {
		head := helper.BuildLinkedList(c.head)
		rs := DoubleIt(head)
		got := helper.BuildLinkedListToArray(rs)

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("DoubleIt test failed, want %v - got %v", c.want, got)
		}
	}
}

func TestDoubleItTwoPointers(t *testing.T) {
	cases := []struct {
		head []int
		want []int
	}{
		{[]int{1, 8, 1}, []int{3, 6, 2}},
		{[]int{1, 8, 9}, []int{3, 7, 8}},
		{[]int{9, 9, 9}, []int{1, 9, 9, 8}},
	}

	for _, c := range cases {
		head := helper.BuildLinkedList(c.head)
		rs := DoubleItTwoPointers(head)
		got := helper.BuildLinkedListToArray(rs)

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("DoubleIt test failed, want %v - got %v", c.want, got)
		}
	}
}
