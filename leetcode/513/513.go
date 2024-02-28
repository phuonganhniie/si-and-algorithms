package leetcode_513

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
	bfsQueue := make([]GraphNode, 0)
	bfsQueue = append(bfsQueue, GraphNode{root, 0})

	var currentDepth = -1
	var leftMost int

	for len(bfsQueue) > 0 {
		front := bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		if front.depth > currentDepth {
			currentDepth = front.depth
			leftMost = front.node.Val
		}

		// enqueue children
		if front.node.Left != nil {
			bfsQueue = append(bfsQueue, GraphNode{front.node.Left, front.depth + 1})
		}

		if front.node.Right != nil {
			bfsQueue = append(bfsQueue, GraphNode{front.node.Right, front.depth + 1})
		}
	}
	return leftMost
}
