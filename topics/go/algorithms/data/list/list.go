// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package list implements of a doubly link list in Go.
package list

import (
	"fmt"
	"strings"
)

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

	// When creating the new node, have the new node
	// point to the last node in the list.
	n := Node{
		Data: data,
		prev: l.last,
	}

	// Increment the count for the new node.
	l.Count++

	// If this is the first node, attach it.
	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	// Fix the fact that the last node does not point back to the NEW node.
	l.last.next = &n

	//              First                                       Last           l.last.next
	//                V                                           V                    V
	// nil <- Prev.[Node0].Next <-> Prev.[Node1].Next <-> Prev.[Node2].Next <-> Prev.[NEW].Next -> nil

	// Fix the fact the Last pointer is not pointing to the true end of the list.
	l.last = &n

	//              First                                       Last                 Last
	//                V                                           V  ----> MOVE ---->  V
	// nil <- Prev.[Node0].Next <-> Prev.[Node1].Next <-> Prev.[Node2].Next <-> Prev.[NEW].Next -> nil

	return &n
}

// AddFront places a new node at the front of the list.
func (l *List) AddFront(data string) *Node {

	// When creating the new node, have the new node
	// point to the first node in the list.
	n := Node{
		Data: data,
		next: l.first,
	}

	// Increment the count for the new node.
	l.Count++

	// If this is the first node, attach it.
	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	// Fix the fact that the first node does not point back to the NEW node.
	l.first.prev = &n

	//      l.first.prev                First                                       Last
	//               V                    V                                           V
	// nil <- Prev.[NEW].Next <-> Prev.[Node2].Next <-> Prev.[Node1].Next <-> Prev.[Node0].Next -> nil

	// Fix the fact the First pointer is not pointing to the true beginning of the list.
	l.first = &n

	//             First                First                                       Last
	//               V  <----> MOVE <---- V                                           V
	// nil <- Prev.[NEW].Next <-> Prev.[Node2].Next <-> Prev.[Node1].Next <-> Prev.[Node0].Next -> nil

	return &n
}

// Find traverses the list looking for the specified data.
func (l *List) Find(data string) (*Node, error) {
	n := l.first
	for n != nil {
		if n.Data == data {
			return n, nil
		}
		n = n.next
	}
	return nil, fmt.Errorf("unable to locate %q in list", data)
}

// FindReverse traverses the list in the opposite direction
// looking for the specified data.
func (l *List) FindReverse(data string) (*Node, error) {
	n := l.last
	for n != nil {
		if n.Data == data {
			return n, nil
		}
		n = n.prev
	}
	return nil, fmt.Errorf("unable to locate %q in list", data)
}

// Remove traverses the list looking for the specified data
// and if found, removes the node from the list.
func (l *List) Remove(data string) (*Node, error) {
	n, err := l.Find(data)
	if err != nil {
		return nil, err
	}

	// Detach the node by linking the previous node's next
	// pointer to the node in front of the one being removed.
	n.prev.next = n.next
	n.next.prev = n.prev
	l.Count--

	return n, nil
}

// Operate accepts a function that takes a node and calls
// the specified function for every node found.
func (l *List) Operate(f func(n *Node) error) error {
	n := l.first
	for n != nil {
		if err := f(n); err != nil {
			return err
		}
		n = n.next
	}
	return nil
}

// OperateReverse accepts a function that takes a node and
// calls the specified function for every node found.
func (l *List) OperateReverse(f func(n *Node) error) error {
	n := l.last
	for n != nil {
		if err := f(n); err != nil {
			return err
		}
		n = n.prev
	}
	return nil
}

// AddSort adds a node based on lexical ordering.
func (l *List) AddSort(data string) *Node {

	// If the list was empty add the data
	// as the first node.
	if l.first == nil {
		return l.Add(data)
	}

	// Traverse the list looking for placement.
	n := l.first
	for n != nil {

		// If this data is greater than the current node,
		// keep traversing until it is less than or equal.
		if strings.Compare(data, n.Data) > 0 {
			n = n.next
			continue
		}

		// Create the new node and place it before the
		// current node.
		new := Node{
			Data: data,
			next: n,
			prev: n.prev,
		}

		l.Count++

		// If this node is now to be the first,
		// fix the first pointer.
		if l.first == n {
			l.first = &new
		}

		// If the current node points to a previous node,
		// then that previous nodes next must point to the
		// new node.
		if n.prev != nil {
			n.prev.next = &new
		}

		// The current previous points must point back
		// to this new node.
		n.prev = &new

		return n
	}

	// This must be the largest string, so add to the end.
	return l.Add(data)
}
