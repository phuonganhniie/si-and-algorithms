package tree

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func inOrderTraversal(root *helper.NodeTree, res *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, res)
	*res = append(*res, root.Value)
	inOrderTraversal(root.Right, res)
}

func TestLevelOrderBinaryTree(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 2, 5, 1, 6, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{8, 4, 9, 2, 10, 5, 1, 6, 3, 7}},
	}

	countErr := 0
	for _, tt := range tests {
		tree := LevelOrderBinaryTree(tt.arr)
		var result []int
		inOrderTraversal(tree.Root, &result)

		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("LevelOrderBinaryTree test failed, want: %v - got: %v", tt.want, result)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
