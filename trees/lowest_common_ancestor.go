package trees

// LCA returns the lowest common ancestor of the
// two nodes
//				 5
//             /  \
//            1    4
//           / \  / \
//          3  8 9  2
//         / \
//        6  7
// examples:
// lca(root,8,7) = 1
// lca(root,1,5) = 5
// lca(root,2,2) = 2
// lca(root,2,0) = nil
// lca(nil ,1,5) = nil
func LCA(root *Node, node1, node2 int) *int {
	if root == nil {
		return nil
	}
	if node1 == node2 {
		return &node1
	}
	if root.Value == node1 || root.Value == node2 {
		return &root.Value
	}
	left := LCA(root.Left, node1, node2)
	right := LCA(root.Right, node1, node2)
	if left != nil && right != nil {
		return &root.Value
	}
	if left == nil {
		return right
	}
	return left
}
