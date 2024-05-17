package leetcode_1325

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestLeafNodes(t *testing.T) {
	cases := []struct {
		root   []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 2, 0, 2, 4}, 2, []int{1, 3, 4}},
		{[]int{1, 3, 3, 3, 2}, 3, []int{1, 3, 2}},
		{[]int{1, 2, 0, 2, 0, 2}, 2, []int{1}},
	}

	for _, c := range cases {
		root := helper.BuildBinaryTree(c.root)
		result := RemoveLeafNodes(root, c.target)
		got := helper.BuildArrayFromBinaryTree(result)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RemoveLeafNodes test failed, want %v - got %v", c.want, got)
		}
	}
}
