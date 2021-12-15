package main

import (
	"fmt"
	"math/big"
)

// =============================================================================

// Defining a type that implements a part of the same API as big.Float. This
// will provide two implementation of the same method set. This API uses
// pointer semantics for everything related to the API.

type Float64 float64

func (f *Float64) Add(a, b *Float64) *Float64 {
	if f == nil {
		f = new(Float64)
	}
	*f = *a + *b
	return f
}

func (f *Float64) Mul(a, b *Float64) *Float64 {
	if f == nil {
		f = new(Float64)
	}
	*f = *a * *b
	return f
}

func (f *Float64) String() string {
	return fmt.Sprintf("%.4f", *f)
}

// ==============================================================================

// The EvalPoly function is using pointer semantics to accept a slice of
// pointers of some type T, a pointer to an individual value of some type
// T and returns a pointer of a new value of some type T.
//
// Since the method set is using pointer semantics, the type list on line
// 48 is necessary. This tells the compiler that PT is a pointer type of T.
// This allows for the conversion syntax to work on line 54 and the compiler
// to accepts the pointer semantic methods for type T.

type Float[T any] interface {
	*T
	Add(*T, *T) *T
	Mul(*T, *T) *T
}

func EvalPoly[T any, PT Float[T]](s []PT, v PT) PT {
	sum := PT(new(T))
	for _, coef := range s {
		sum.Mul(sum, v)
		sum.Add(sum, coef)
	}
	return sum
}

// ==============================================================================

// This function is a convenience function for getting the address for a literal
// floating point number.

func ref(f Float64) *Float64 {
	return &f
}

// On the calls to EvalPoly, the type information for the value and pointer type
// need to be explicitly passed. The draft suggests that in the future, just the
// type name will need to be provided.

func main() {	
	slice1 := []*big.Float{big.NewFloat(3.12), big.NewFloat(5.4), big.NewFloat(1.7)}
	f1 := EvalPoly[big.Float, *big.Float](slice1, big.NewFloat(3.1))
	fmt.Println(f1.String())
	
	slice2 := []*Float64{ref(3.12), ref(5.4), ref(1.7)}
	f2 := EvalPoly[Float64, *Float64](slice2, ref(3.1))
	fmt.Println(f2.String())
}
