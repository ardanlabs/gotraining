package main

import (
	"fmt"
)

// =============================================================================

// This code defines two user defined types that implement a linked list. The
// node type contains data of some type T (to be determined later) and points
// to other nodes of the same type T. The list type contains pointers to the
// first and last nodes of some type T. The add method is declared with a
// pointer receiver based on a list of some type T and is implemented to add
// nodes to the list of the same type T.

type node[T any] struct {
	Data T
	next *node[T]
	prev *node[T]
}

type list[T any] struct {
	first *node[T]
	last  *node[T]
}

func (l *list[T]) add(data T) *node[T] {
	n := node[T]{
		Data: data,
		prev: l.last,
	}
	if l.first == nil {
		l.first = &n
		l.last = &n
		return &n
	}
	l.last.next = &n
	l.last = &n
	return &n
}

// =============================================================================

// This user type represents the data to be stored into the linked list.

type user struct {
	name string
}

// =============================================================================

func main() {

	// Store values of type user into the list.
	var lv list[user]
	n1 := lv.add(user{"bill"})
	n2 := lv.add(user{"ale"})
	fmt.Println(n1.Data, n2.Data)

	// Store pointers of type user into the list.
	var lp list[*user]
	n3 := lp.add(&user{"bill"})
	n4 := lp.add(&user{"ale"})
	fmt.Println(n3.Data, n4.Data)
}
