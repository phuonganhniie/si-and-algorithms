package leetcode_2487

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode"
)

func TestRemoveNodes(t *testing.T) {
	cases := []struct {
		initial []int
		want    []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{13, 8}},
		{[]int{5, 2, 13, 3, 8}, []int{13, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for _, c := range cases {
		head := leetcode.BuildLinkedList(c.initial)
		got := RemoveNodes(head)
		actual := leetcode.BuildLinkedListToArray(got)
		if !reflect.DeepEqual(c.want, actual) {
			t.Errorf("RemoveNodes failed, want %v - got %v", c.want, actual)
		}
	}
}
