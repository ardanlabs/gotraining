package main

// Tree represents all values in the tree.
type Tree struct {
	root *Node
}

// Insert adds a value into the tree.
func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}

	t.root.insert(value)
}

// Node represents a value in the tree.
type Node struct {
	value int
	left  *Node
	right *Node
}

// insert adds the value into the tree.
func (n *Node) insert(value int) {
	switch {
	case value <= n.value:
		if n.left == nil {
			n.left = &Node{value: value}
			return
		}
		n.left.insert(value)
	case value > n.value:
		if n.right == nil {
			n.right = &Node{value: value}
			return
		}
		n.right.insert(value)
	}
}

func main() {
	values := []int{40, 5, 10, 80, 62, 2, 45, 12, 23, 77, 3, 2}

	var t Tree
	for _, value := range values {
		t.Insert(value)
	}

	PrettyPrint(&t)
}
