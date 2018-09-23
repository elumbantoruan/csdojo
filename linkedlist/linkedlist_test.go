package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NthFromEnd return nth element value from the end of the list
// list = [1 2 3]
//    n = 1
//  val = 3
//
// list = [5 4 3 2 1]
//    n = 3
//  val = 3
//
// list = [5 4 3 2 1]
//    n = 6
//  val = nil
func TestNthFromEnd(t *testing.T) {
	type args struct {
		head *Node
		n    int
	}
	var (
		exp = 3
	)
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "test1",
			args: args{head: createList1(), n: 1},
			want: &exp,
		},
		{
			name: "test2",
			args: args{head: createList2(), n: 3},
			want: &exp,
		},
		{
			name: "test3",
			args: args{head: createList2(), n: 6},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NthFromEnd(tt.args.head, tt.args.n)
			assert.Equal(t, tt.want, result)
		})
	}
}

func createList1() *Node {
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}

	return head
}

func createList2() *Node {
	head := &Node{Value: 5}
	head.Next = &Node{Value: 4}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 2}
	head.Next.Next.Next.Next = &Node{Value: 1}
	return head
}
