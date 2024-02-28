package leetcode_513

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type GraphNode struct {
	node  *TreeNode
	depth int
}

func FindBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	bfsQueue := make([]GraphNode, 0)
	bfsQueue = append(bfsQueue, GraphNode{root, 0})

	var currentDepth = -1
	var leftMost int

	for len(bfsQueue) > 0 {
		front := bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		fmt.Printf("process node: %d at depth: %d\n", front.node.Val, front.depth)

		if front.depth > currentDepth {
			currentDepth = front.depth
			leftMost = front.node.Val
			fmt.Printf("new leftMost: %d at depth: %d\n", leftMost, currentDepth)
		}

		// enqueue children
		if front.node.Left != nil {
			fmt.Printf("enqueue left child of node %d\n", front.node.Val)
			bfsQueue = append(bfsQueue, GraphNode{front.node.Left, front.depth + 1})
		}

		if front.node.Right != nil {
			fmt.Printf("enqueue right child of node %d\n", front.node.Val)
			bfsQueue = append(bfsQueue, GraphNode{front.node.Right, front.depth + 1})
		}
	}
	return leftMost
}
