package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	rs := []float64{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSum := 0
		levelCount := len(queue)

		for i := 0; i < levelCount; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		avg := float64(levelSum) / float64(levelCount)
		rs = append(rs, avg)
	}

	return rs
}
