package binary

// InOrder traversal travel from the leftmost node to the rightmost nodes
// regardless of depth.
// In-order traversal gives node values in ascending order.
//
//          #4
//       /      \
//      #2      #6
//     /  \    /  \
//    #1  #3  #5  #7
//
func InOrder(tree *Tree) []int {
	order := []int{}
	f := func(n *node) {
		order = append(order, n.value)
	}
	inOrder(tree.root, f)
	return order
}

func inOrder(node *node, f func(*node)) {
	if node != nil {
		inOrder(node.left, f)
		f(node)
		inOrder(node.right, f)
	}
}

// PreOrder traversal get the root note then traversing its child
// nodes recursively.
// Use cases: copying tree, mapping prefix notation.
//          #1
//       /      \
//      #2      #5
//     /  \    /  \
//    #3  #4  #6  #7
//
func PreOrder(tree *Tree) []int {
	order := []int{}
	preOrder(tree.root, func(n *node) { order = append(order, n.value) })
	return order
}

func preOrder(node *node, f func(*node)) {
	if node != nil {
		f(node)
		preOrder(node.left, f)
		preOrder(node.right, f)
	}

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
//
func PostOrder(tree *Tree) []int {
	order := []int{}
	postOrder(tree.root, func(n *node) { order = append(order, n.value) })
	return order
}

// Node value is recorded after left child note and then right child node
func postOrder(node *node, f func(*node)) {
	if node != nil {
		postOrder(node.left, f)
		postOrder(node.right, f)
		f(node)
	}
}
