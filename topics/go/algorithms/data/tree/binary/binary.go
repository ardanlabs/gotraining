// Package binary is an implementation of a binary tree.
package binary

// Tree represents all values in the tree.
type Tree struct {
	root *node
}

// Insert adds a value into the tree.
func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &node{value: value}
		return
	}

	t.root.Insert(value)
}

type node struct {
	value int
	left  *node
	right *node
}

// insert adds the value into the tree.
func (n *node) Insert(value int) {
	switch {
	case value <= n.value:
		if n.left == nil {
			n.left = &node{value: value}
			return
		}
		n.left.Insert(value)
	case value > n.value:
		if n.right == nil {
			n.right = &node{value: value}
			return
		}
		n.right.Insert(value)
	}
}
