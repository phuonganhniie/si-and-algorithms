package dfs

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		fmt.Printf("Reached a null node, returning false \n")
		return false
	}

	// check if it's a leaf node
	newTargetSum := targetSum - root.Val
	fmt.Printf("Visiting Node with Value: %d | Current Target Sum: %d \n", root.Val, newTargetSum)

	if root.Left == nil && root.Right == nil {
		if newTargetSum == 0 {
			fmt.Printf("Leaf node with Value: %d | Path sum equals targetSum, returning true \n", root.Val)
			return true
		} else {
			fmt.Printf("Leaf node with Value: %d | Path sum does not equal targetSum, returning false \n", root.Val)
			return false
		}
	}

	leftCheck := false
	if root.Left != nil {
		leftCheck = HasPathSum(root.Left, newTargetSum)
	}

	rightCheck := false
	if root.Right != nil {
		rightCheck = HasPathSum(root.Right, newTargetSum)
	}

	return leftCheck || rightCheck
}
