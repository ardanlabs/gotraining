// Implementation of tree traversal algorithms of binary.
// Given a binary tree, the functions return the traversing order slice.
// The traversal functions have internal function which traverse nodes
// recursively.
// Closure is used in the internal function to capture the node values.
package binary

// In-order traversal travel from the leftmost node to the rightmost nodes
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
	inOrder(tree.Root, func(n *Node) { order = append(order, n.Value) })
	return order
}

// Node value is recorded after the left child but before the right child.
func inOrder(node *Node, f func(*Node)) {
	if node != nil {
		inOrder(node.Left, f)
		f(node)
		inOrder(node.Right, f)
	}
}

// Pre-order traversal get the root note then traversing its child
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
	preOrder(tree.Root, func(n *Node) { order = append(order, n.Value) })
	return order
}

// Node value is recorded first followed by left child then right child.
func preOrder(node *Node, f func(*Node)) {
	if node != nil {
		f(node)
		preOrder(node.Left, f)
		preOrder(node.Right, f)
	}

}

// Post-order traversal get the leftmost node then its sibling then go up to its
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
	postOrder(tree.Root, func(n *Node) { order = append(order, n.Value) })
	return order
}

// Node value is recorded after left child note and then right child node
func postOrder(node *Node, f func(*Node)) {
	if node != nil {
		postOrder(node.Left, f)
		postOrder(node.Right, f)
		f(node)
	}
}
