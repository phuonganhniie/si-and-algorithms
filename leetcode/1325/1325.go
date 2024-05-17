package leetcode_1325

import (
	"github.com/phuonganhniie/leetcode/helper"
)

func RemoveLeafNodes(root *helper.TreeNode, target int) *helper.TreeNode {
	if root == nil {
		return nil
	}

	// 1. Visit the left children
	root.Left = RemoveLeafNodes(root.Left, target)

	// 2. Visit the right children
	root.Right = RemoveLeafNodes(root.Right, target)

	// 3. Check if the current node is a leaf and its value equals target
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	// 4. Keep it untouched otherwise
	return root
}
