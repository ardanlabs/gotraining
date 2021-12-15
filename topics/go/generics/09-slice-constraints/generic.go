package main

import (
	"fmt"
)

// =============================================================================

// operateFunc defines a function type that takes a value of some type T and
// returns a value of the same type T (to be determined later).
//
// Slice defines a constraint that the data is a slice of some type T (to be
// determined later).

type operateFunc[T any] func(t T) T

type Slice[T any] interface {
	~ []T
}

// When it's important that the slice being passed in is exactly the same
// as the slice being returned, use a slice contraint. This ensures that the
// result slice S is the same as the incoming slice S.

func operate[S Slice[T], T any](slice S, fn operateFunc[T]) S {
	ret := make(S, len(slice))
	for i, v := range slice {
		ret[i] = fn(v)
	}
	return ret
}

// If you don't care about the constraint defined above, then operate2 provides
// a simpler form. Operate2 still works because you can assign a slice of some
// type T to the input and output arguments. However, the concrete types of the
// input and output arguments will be based on the underlying types. In this
// case not a slice of Numbers, but a slice of integers. This is not the case
// with operate function above.

func operate2[T any](slice []T, fn operateFunc[T]) []T {
	ret := make([]T, len(slice))
	for i, v := range slice {
		ret[i] = fn(v)
	}
	return ret
}

// I suspect most of the time operate2 is what you want: it's simpler, and more
// flexible: You can always assign a []int back to a Numbers variable and vice
// versa. But if you need to preserve that incoming type in the result for some
// reason, you will need to use operate.

// =============================================================================

// This code defines a named type that is based on a slice of integers. An
// integer is the underlying type.
//
// Double is a function that takes a value of type Numbers, multiplies each
// value in the underlying integer slice and returns that new Numbers value.
//
// Line 73 is commented out because the compiler is smart enough to infer the
// types for S and T. The commented code shows the types being infered.
//
// operate2 is not used in the example.

type Numbers []int

func DoubleUnderlying(n Numbers) Numbers {
	fn := func(n int) int {
		return 2 * n
	}

	numbers := operate2(n, fn)
	fmt.Printf("%T", numbers)
	return numbers
}

func DoubleUserDefined(n Numbers) Numbers {
	fn := func(n int) int {
		return 2 * n
	}

	numbers := operate(n, fn)
	fmt.Printf("%T", numbers)
	return numbers
}

// =============================================================================

func main() {
	n := Numbers{10, 20, 30, 40, 50}
	fmt.Println(n)

	n = DoubleUnderlying(n)
	fmt.Println(n)

	n = DoubleUserDefined(n)
	fmt.Println(n)
}