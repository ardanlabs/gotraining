package main

import (
	"fmt"
)

// =============================================================================

// Print is a function that accepts a slice of some type L and a slice of some
// type V (to be determined later). Each value of the labels slice will be joined
// with the vals slice at the same index position. This code shows how the
// generics type list can contain more than just one generic type and have
// different constraints for each.

func Print[L any, V fmt.Stringer](labels []L, vals []V) {
	for i, v := range vals {
		fmt.Println(labels[i], v.String())
	}
}

// =============================================================================

// This code defines a concrete type named user that implements the fmt.Stringer
// interface. The String method just returns the name field from the user type.

type user struct {
	name string
}

func (u user) String() string {
	return u.name
}

// =============================================================================

func main() {
	labels := []int{1, 2, 3}
	names := []user{{"bill"},{"jill"},{"joan"}}
	Print(labels, names)
}