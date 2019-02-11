package list

// Node represents the data being stored.
type Node struct {
	Data string
	next *Node
	prev *Node
}

// List represents a list of nodes.
type List struct {
	Count int
	first *Node
	last  *Node
}

// AddOk uses the `newNode` variable for the node. This variable
// name gets in the way of readability. It becomes the focal
// point when trying to read the code.
func (l *List) AddOk(data string) *Node {
	newNode := Node{
		Data: data,
		prev: l.last,
	}

	switch l.Count {
	case 0:
		l.first = &newNode
		l.last = &newNode
	default:
		l.last.next = &newNode
		l.last = &newNode
	}

	l.Count++
	return &newNode
}

// AddBetter uses the `n` variable for the node. This variable is better
// because it doesn't get in the way of readability. Now the switch/case
// logic becomes the focal point when reading the code.
func (l *List) AddBetter(data string) *Node {
	n := Node{
		Data: data,
		prev: l.last,
	}

	switch l.Count {
	case 0:
		l.first = &n
		l.last = &n
	default:
		l.last.next = &n
		l.last = &n
	}

	l.Count++
	return &n
}
