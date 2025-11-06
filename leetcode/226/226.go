package leetcode_226

type TreeNode struct {
	val       int
	leftNode  *TreeNode
	rightNode *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.leftNode, root.rightNode = root.rightNode, root.leftNode
	invertTree(root.leftNode)
	invertTree(root.rightNode)

	return root
}
