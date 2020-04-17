package bst

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
func (bst *BST) Max() int {

	currentNode := bst.root

	if currentNode == nil {
		return -1
	}

	for {
		if currentNode.right == nil {
			return currentNode.Data
		}
		currentNode = currentNode.right
	}
}

// Min returns the min int
func (bst *BST) Min() int {

	currentNode := bst.root

	if currentNode == nil {
		return -1
	}

	for {
		if currentNode.left == nil {
			return currentNode.Data
		}
		currentNode = currentNode.left
	}
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
