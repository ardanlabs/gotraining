package bst

import "fmt"

// TODO This is currently a draft...

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

// New creates a new BST.
func New() *BST {
	return &BST{
		root: nil,
	}
}

// Insert an integer into the BST.
func (bst *BST) Insert(value int) {

	// Create a new node.
	n := &Node{
		Data:  value,
		left:  nil,
		right: nil,
	}

	// If the root node is nil set the root node
	// to n
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// Max returns the max int
func (bst *BST) Max() (int, error) {

	currentNode := bst.root

	if currentNode == nil {
		return 0, fmt.Errorf("root node is nil")
	}

	for {
		if currentNode.right == nil {
			return currentNode.Data, nil
		}
		currentNode = currentNode.right
	}
}

// Min returns the min int
func (bst *BST) Min() (int, error) {

	currentNode := bst.root

	if currentNode == nil {
		return 0, fmt.Errorf("root node is nil")
	}

	for {
		if currentNode.left == nil {
			return currentNode.Data, nil
		}
		currentNode = currentNode.left
	}
}

// Search return a bool.
func (bst *BST) Search(value int) bool {

	currentNode := bst.root

	if currentNode == nil {
		return false
	}

	if value == currentNode.Data {
		return true
	}

	return false
}

func insertNode(root, newNode *Node) {

	// Insert into the left side of the tree.
	if newNode.Data < root.Data {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}

		// Insert into the right side of the tree.
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}
