package bfs

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		minDepth++
		levelCount := len(queue)

		for i := 0; i < levelCount; i++ {
			node := queue[0]
			queue = queue[1:]

			// check if the node is a leaf
			if node.Left == nil && node.Right == nil {
				return minDepth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return minDepth
}
