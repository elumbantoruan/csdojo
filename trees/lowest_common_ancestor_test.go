package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
func TestLCA(t *testing.T) {
	type args struct {
		root  *Node
		node1 int
		node2 int
	}
	want1 := 1
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "Test1",
			args: args{
				root:  createTree(),
				node1: 8,
				node2: 7,
			},
			want: &want1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LCA(tt.args.root, tt.args.node1, tt.args.node2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func createTree() *Node {
	root := Node{Value: 5}
	root.Left = &Node{Value: 1}
	root.Right = &Node{Value: 4}

	root.Left.Left = &Node{Value: 3}
	root.Left.Right = &Node{Value: 8}

	root.Right.Left = &Node{Value: 9}
	root.Right.Right = &Node{Value: 2}

	root.Left.Left.Left = &Node{Value: 6}
	root.Left.Left.Right = &Node{Value: 7}

	return &root
}
