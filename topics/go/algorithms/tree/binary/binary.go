package binary

// Tree represents all values in the tree.
type Tree struct {
	Root *Node
}

// Insert adds a value into the tree.
func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}

	t.Root.Insert(value)
}

// Node represents a value in the tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// insert adds the value into the tree.
func (n *Node) Insert(value int) {
	switch {
	case value <= n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return
		}
		n.Left.Insert(value)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return
		}
		n.Right.Insert(value)
	}
}
