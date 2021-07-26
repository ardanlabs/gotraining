// Package binary is an implementation of a balanced binary tree.
package binary

// Tree represents all values in the tree.
type Tree struct {
	root *node
}

// Insert adds a value into the tree and keeps the tree balanced.
func (t *Tree) Insert(value int) {
	t.root = t.root.insert(t, value)

	if t.root.balRatio() < -1 || t.root.balRatio() > 1 {
		t.root = t.root.rebalance()
	}
}

// PreOrder traversal get the root node then traversing its child
// nodes recursively.
// Use cases: copying tree, mapping prefix notation.
//          #1
//       /      \
//      #2      #5
//     /  \    /  \
//    #3  #4  #6  #7
func (t *Tree) PreOrder() []int {
	order := []int{}
	f := func(n *node) {
		order = append(order, n.Value)
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
func (t *Tree) InOrder() []int {
	order := []int{}
	f := func(n *node) {
		order = append(order, n.Value)
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
func (t *Tree) PostOrder() []int {
	order := []int{}
	f := func(n *node) {
		order = append(order, n.Value)
	}
	t.root.postOrder(f)
	return order
}

// =============================================================================

// node represents the data stored in the tree.
type node struct {
	Value int
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
func (n *node) insert(t *Tree, value int) *node {
	if n == nil {
		return &node{Value: value, level: 1, tree: t}
	}

	switch {
	case value < n.Value:
		n.left = n.left.insert(t, value)
	case value > n.Value:
		n.right = n.right.insert(t, value)
	default:
		return n.rebalance()
	}
	n.level = max(n.left.height(), n.right.height()) + 1

	return n.rebalance()
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
