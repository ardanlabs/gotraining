// Package binary is an implementation of a balanced binary tree.
// A majority of this code comes from appliedgo.
// https://appliedgo.net/balancedtree/
// https://appliedgo.net/bintree/
package binary

import (
	"errors"
)

// Data represents the information being stored.
type Data struct {
	Key  int
	Name string
}

// Tree represents all values in the tree.
type Tree struct {
	root *node
}

// Insert adds a value into the tree and keeps the tree balanced.
func (t *Tree) Insert(data Data) {
	t.root = t.root.insert(t, data)

	if t.root.balRatio() < -1 || t.root.balRatio() > 1 {
		t.root = t.root.rebalance()
	}
}

// Find traverses the tree looking for the specified tree.
func (t *Tree) Find(key int) (Data, error) {
	if t.root == nil {
		return Data{}, errors.New("cannot find from an empty tree")
	}

	return t.root.find(key)
}

// Delete removes the key from the tree and keeps it balaned.
func (t *Tree) Delete(key int) error {
	if t.root == nil {
		return errors.New("cannot delete from an empty tree")
	}

	fakeParent := &node{right: t.root}
	if err := t.root.delete(key, fakeParent); err != nil {
		return err
	}

	if fakeParent.right == nil {
		t.root = nil
	}
	return nil
}

// PreOrder traversal get the root node then traversing its child
// nodes recursively.
// Use cases: copying tree, mapping prefix notation.
//          #1
//       /      \
//      #2      #5
//     /  \    /  \
//    #3  #4  #6  #7
func (t *Tree) PreOrder() []Data {
	order := []Data{}
	f := func(n *node) {
		order = append(order, n.data)
	}
	t.root.preOrder(f)
	return order
}

// InOrder traversal travel from the leftmost node to the rightmost nodes
// regardless of depth.
// In-order traversal gives node values in ascending order.
//
//          #4
//       /      \
//      #2      #6
//     /  \    /  \
//    #1  #3  #5  #7
func (t *Tree) InOrder() []Data {
	order := []Data{}
	f := func(n *node) {
		order = append(order, n.data)
	}
	t.root.inOrder(f)
	return order
}

// PostOrder traversal get the leftmost node then its sibling then go up to its
// parent, recursively.
// Use cases: tree deletion, mapping postfix notation.
//
//          #7
//       /      \
//      #3      #6
//     /  \    /  \
//    #1  #2  #4  #5
func (t *Tree) PostOrder() []Data {
	order := []Data{}
	f := func(n *node) {
		order = append(order, n.data)
	}
	t.root.postOrder(f)
	return order
}

// =============================================================================

// node represents the data stored in the tree.
type node struct {
	data  Data
	level int
	tree  *Tree
	left  *node
	right *node
}

// height returned the level of the tree the node exists in.
// Level 1 is at the last layer of the tree.
//
//          #7          -- height = 3
//       /      \
//      #3      #6      -- height = 2
//     /  \    /  \
//    #1  #2  #4  #5    -- height = 1
func (n *node) height() int {
	if n == nil {
		return 0
	}
	return n.level
}

// insert adds the node into the tree and makes sure the
// tree stays balanced.
func (n *node) insert(t *Tree, data Data) *node {
	if n == nil {
		return &node{data: data, level: 1, tree: t}
	}

	switch {
	case data.Key < n.data.Key:
		n.left = n.left.insert(t, data)

	case data.Key > n.data.Key:
		n.right = n.right.insert(t, data)

	default:
		return n.rebalance()
	}

	n.level = max(n.left.height(), n.right.height()) + 1
	return n.rebalance()
}

// find traverses the tree looking for the specified key.
func (n *node) find(key int) (Data, error) {
	if n == nil {
		return Data{}, errors.New("key not found")
	}

	switch {
	case n.data.Key == key:
		return n.data, nil

	case key < n.data.Key:
		return n.left.find(key)

	default:
		return n.right.find(key)
	}
}

// balRatio provides information about the balance ratio
// of the node.
func (n *node) balRatio() int {
	return n.right.height() - n.left.height()
}

// rotateLeft turns the node to the left.
//
//   #3          #4
//     \        /  \
//     #4      #3  #5
//       \
//       #5
func (n *node) rotateLeft() *node {
	r := n.right
	n.right = r.left
	r.left = n
	n.level = max(n.left.height(), n.right.height()) + 1
	r.level = max(r.left.height(), r.right.height()) + 1
	return r
}

// rotateRight turns the node to the right.
//
//       #5      #4
//      /       /  \
//     #4      #3  #5
//    /
//   #3
func (n *node) rotateRight() *node {
	l := n.left
	n.left = l.right
	l.right = n
	n.level = max(n.left.height(), n.right.height()) + 1
	l.level = max(l.left.height(), l.right.height()) + 1
	return l
}

// rotateLeftRight turns the node to the left and then right.
//
//     #5          #5      #4
//    /           /       /  \
//   #3          #4      #3  #5
//     \        /
//     #4      #3
func (n *node) rotateLeftRight() *node {
	n.left = n.left.rotateLeft()
	n = n.rotateRight()
	n.level = max(n.left.height(), n.right.height()) + 1
	return n
}

// rotateLeftRight turns the node to the left and then right.
//
//   #3        #3          #4
//     \         \        /  \
//     #5        #4      #3  #5
//    /            \
//   #4            #5
func (n *node) rotateRightLeft() *node {
	n.right = n.right.rotateRight()
	n = n.rotateLeft()
	n.level = max(n.left.height(), n.right.height()) + 1
	return n
}

// rebalance will rotate the nodes based on the ratio.
func (n *node) rebalance() *node {
	switch {
	case n.balRatio() < -1 && n.left.balRatio() == -1:
		return n.rotateRight()

	case n.balRatio() > 1 && n.right.balRatio() == 1:
		return n.rotateLeft()

	case n.balRatio() < -1 && n.left.balRatio() == 1:
		return n.rotateLeftRight()

	case n.balRatio() > 1 && n.right.balRatio() == -1:
		return n.rotateRightLeft()
	}
	return n
}

// findMax finds the maximum element in a (sub-)tree. Its value replaces
// the value of the to-be-deleted node. Return values: the node itself and
// its parent node.
func (n *node) findMax(parent *node) (*node, *node) {
	switch {
	case n == nil:
		return nil, parent

	case n.right == nil:
		return n, parent
	}
	return n.right.findMax(n)
}

// replaceNode replaces the parent’s child pointer to n with a pointer
// to the replacement node. parent must not be nil.
func (n *node) replaceNode(parent, replacement *node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	switch n {
	case parent.left:
		parent.left = replacement

	default:
		parent.right = replacement
	}

	return nil
}

// delete removes an element from the tree. It is an error to try
// deleting an element that does not exist. In order to remove an
// element properly, Delete needs to know the node’s parent node.
// Parent must not be nil.
func (n *node) delete(key int, parent *node) error {
	if n == nil {
		return errors.New("value to be deleted does not exist in the tree")
	}

	switch {
	case key < n.data.Key:
		return n.left.delete(key, n)

	case key > n.data.Key:
		return n.right.delete(key, n)

	default:
		switch {
		case n.left == nil && n.right == nil:
			n.replaceNode(parent, nil)
			return nil
		case n.left == nil:
			n.replaceNode(parent, n.right)
			return nil
		case n.right == nil:
			n.replaceNode(parent, n.left)
			return nil
		}
		replacement, replParent := n.left.findMax(n)
		n.data = replacement.data
		return replacement.delete(replacement.data.Key, replParent)
	}
}

// preOrder traverses the node by traversing the child nodes recursively.
func (n *node) preOrder(f func(*node)) {
	if n != nil {
		f(n)
		n.left.preOrder(f)
		n.right.preOrder(f)
	}
}

// inOrder traversal the node by the leftmost node to the rightmost nodes
// regardless of depth.
func (n *node) inOrder(f func(*node)) {
	if n != nil {
		n.left.inOrder(f)
		f(n)
		n.right.inOrder(f)
	}
}

// postOrder traversal the node by the leftmost node then its sibling
// then up to its parent, recursively.
func (n *node) postOrder(f func(*node)) {
	if n != nil {
		n.left.postOrder(f)
		n.right.postOrder(f)
		f(n)
	}
}

// =============================================================================

// max returns the larger of the two values.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
