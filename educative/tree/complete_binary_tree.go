package tree

import (
	"github.com/phuonganhniie/educative/helper"
)

func LevelOrderBinaryTree(arr []int) *helper.Tree {
	tree := new(helper.Tree)
	tree.Root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, i, size int) *helper.NodeTree {
	curr := &helper.NodeTree{Value: arr[i], Left: nil, Right: nil}
	left := 2*i + 1
	right := 2*i + 2

	if left < size {
		curr.Left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.Right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}
