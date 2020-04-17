package bst

// Node represents the data being stored.
type Node struct {
	Data  int
	left  *Node
	right *Node
}

// BST represents a binary search tree.
type BST struct {
	root *Node
}
