// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package list asks the student to implement a double link list in Go.
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

// Add places a new node at the end of the list.
func (l *List) Add(data string) *Node {
	return nil
}

// AddFront places a new node at the front of the list.
func (l *List) AddFront(data string) *Node {
	return nil
}

// Find traverses the list looking for the specified data.
func (l *List) Find(data string) (*Node, error) {
	return nil, nil
}

// FindReverse traverses the list in the opposite direction
// looking for the specified data.
func (l *List) FindReverse(data string) (*Node, error) {
	return nil, nil
}

// Remove traverses the list looking for the specified data
// and if found, removes the node from the list.
func (l *List) Remove(data string) (*Node, error) {
	return nil, nil
}

// Operate accepts a function that takes a node and calls
// the specified function for every node found.
func (l *List) Operate(f func(n *Node) error) error {
	return nil
}

// OperateReverse accepts a function that takes a node and
// calls the specified function for every node found.
func (l *List) OperateReverse(f func(n *Node) error) error {
	return nil
}

// AddSort adds a node based on lexical ordering.
func (l *List) AddSort(data string) *Node {
	return nil
}
