package helper

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}

	i := 1
	for i < len(arr) {
		current := queue[0]
		queue = queue[1:]

		if i < len(arr) {
			if arr[i] != 0 {
				current.Left = &TreeNode{Val: arr[i]}
				queue = append(queue, current.Left)
			}
		}
		i++

		if i < len(arr) {
			if arr[i] != 0 {
				current.Right = &TreeNode{Val: arr[i]}
				queue = append(queue, current.Right)
			}
		}
		i++
	}

	return root
}

func BuildArrayFromBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}
