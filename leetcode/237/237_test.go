package leetcode_237

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestDeleteNode(t *testing.T) {
	cases := []struct {
		initial []int
		nodeVal int
		want    []int
	}{
		{[]int{4, 5, 1, 9}, 5, []int{4, 1, 9}},
		{[]int{4, 5, 1, 9}, 1, []int{4, 5, 9}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
	}

	for _, c := range cases {
		head := helper.BuildLinkedList(c.initial)
		node := helper.FindNode(head, c.nodeVal)
		DeleteNode(node)
		actual := helper.BuildLinkedListToArray(head)

		if !reflect.DeepEqual(actual, c.want) {
			t.Errorf("DeleteNode test case failed. Want %v - Got %v", c.want, actual)
		}
	}
}
