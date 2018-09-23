package trees

import "math"

// Node represents tree's node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// IsBinarySearchTree determines if
func IsBinarySearchTree(tree *Node) bool {
	return isBinarySearchTree(tree, math.MinInt32, math.MaxInt32)
}

func isBinarySearchTree(tree *Node, low, high int) bool {

	if tree == nil {
		return true
	}

	if tree.Value < low ||
		tree.Value > high {
		return false
	}

	return isBinarySearchTree(tree.Left, low, tree.Value) &&
		isBinarySearchTree(tree.Right, tree.Value, high)
}
