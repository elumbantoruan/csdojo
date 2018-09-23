package linkedlist

// Node represents a node in a linked list
type Node struct {
	Value int
	Next  *Node
}

// New creates a node's linkedlist
func New() *Node {
	return &Node{}
}

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
func NthFromEnd(head *Node, n int) (value *int) {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
		if p1 == nil {
			return nil
		}
	}
	for p1.Next != nil {
		p2 = p2.Next
		p1 = p1.Next

	}

	p2 = p2.Next

	return &p2.Value
}
