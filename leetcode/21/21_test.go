package leetcode_21

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
	}

	for _, c := range cases {
		list1 := helper.BuildLinkedList(c.list1)
		list2 := helper.BuildLinkedList(c.list2)
		rs := MergeTwoLists(list1, list2)
		got := helper.BuildLinkedListToArray(rs)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeTwoLists test failed, want %v - got %v", c.want, got)
		}
	}
}
