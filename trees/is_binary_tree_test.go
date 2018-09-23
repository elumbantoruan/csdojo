package trees

import (
	"math"
	"testing"
)

func Test_isBinarySearchTree(t *testing.T) {
	type args struct {
		tree *Node
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				tree: createTree1(),
				low:  math.MinInt32,
				high: math.MaxInt32,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				tree: createTree2(),
				low:  math.MinInt32,
				high: math.MaxInt32,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				tree: createTree3(),
				low:  math.MinInt32,
				high: math.MaxInt32,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBinarySearchTree(tt.args.tree, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("isBinarySearchTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

//        3
//      /  \
//     1    5
//    / \  / \
//   0  2 4   6
func createTree1() *Node {
	tree := &Node{Value: 3}
	tree.Left = &Node{Value: 1}
	tree.Right = &Node{Value: 5}

	tree.Left.Left = &Node{Value: 0}
	tree.Left.Right = &Node{Value: 2}

	tree.Right.Left = &Node{Value: 4}
	tree.Right.Right = &Node{Value: 6}

	return tree
}

//        4
//       / \
//      1   2
//     / \
//    0  3
func createTree2() *Node {
	tree := &Node{Value: 4}
	tree.Left = &Node{Value: 1}
	tree.Right = &Node{Value: 2}

	tree.Left.Left = &Node{Value: 0}
	tree.Left.Right = &Node{Value: 3}

	return tree
}

//        4
//      /  \
//     1    5
//    / \  / \
//   0  2 3   6
func createTree3() *Node {
	tree := &Node{Value: 4}
	tree.Left = &Node{Value: 1}
	tree.Right = &Node{Value: 5}

	tree.Left.Left = &Node{Value: 0}
	tree.Left.Right = &Node{Value: 2}

	tree.Right.Left = &Node{Value: 3}
	tree.Right.Right = &Node{Value: 6}

	return tree
}
