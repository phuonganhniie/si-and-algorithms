package helper

type NodeTree struct {
	Value int
	Left  *NodeTree
	Right *NodeTree
}

type Tree struct {
	Root *NodeTree
}
