// This code provided by Sathish VJ.
package main

import "fmt"

// =============================================================================

// These functions are concrete versions of a swap function. For each type of
// data to be swapped, a new function needs to be implemented.

func swapInteger(v1 int, v2 int) (int, int) {
	v2, v1 = v1, v2
	return v1, v2
}

func swapString(v1 string, v2 string) (string, string) {
	v2, v1 = v1, v2
	return v1, v2
}

// =============================================================================

// This function provides an empty interface solution to perform a generic swap.
// This will allow data of any type to be swapped. The data doesn't need to be
// the same in either interface.

func swapInterface(v1 interface{}, v2 interface{}) (interface{}, interface{}) {
	v2, v1 = v1, v2
	return v1, v2
}

// =============================================================================

// This function provides a generics solution to perform a generic swap. This
// implementation has the advanatge that concrete types are being used and only
// data of the same type can be swapped. The caller doesn't require type assertions.

func swap[T any](v1 T, v2 T) (T, T) {
	v2, v1 = v1, v2
	return v1, v2
}

// =============================================================================

func main() {
	n1, n2 := 10, 90
	n1, n2 = swapInteger(n1, n2)
	fmt.Println("swapInteger 10, 90 ->", n1, n2)

	s1, s2 := "hello", "goodbye"
	s1, s2 = swapString(s1, s2)
	fmt.Println("swapString hello, goodbye ->", s1, s2)

	// To apply the swap back to the float variables,
	// a type assertion is required on line 59.
	f1, f2 := 10.2, 90.4
	i1, i2 := swapInterface(f1, f2)
	f1, f2 = i1.(float64), i2.(float64)
	fmt.Println("swapInterface 10.2, 90.4 ->", f1, f2)

	n1, n2 = swap(n1, n2)
	s1, s2 = swap(s1, s2)
	f1, f2 = swap(f1, f2)
	fmt.Println("swap 90, 10 ->", n1, n2)
	fmt.Println("swap goodbye, hello ->", s1, s2)
	fmt.Println("swap 90.4, 10.2 ->", f1, f2)
}