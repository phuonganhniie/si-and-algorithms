package leetcode_100

/**
 * 100. Same Tree
 * https://leetcode.com/problems/same-tree/?envType=daily-question&envId=2024-02-26
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}

	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		queue = append(queue, node1.Left, node2.Left)
		queue = append(queue, node1.Right, node2.Right)
	}
	return true
}
