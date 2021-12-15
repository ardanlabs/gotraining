package main

import (
	"fmt"
	"math"
)

// ==============================================================================

// The Min function needs to know what type `T` becomes so it can perform a
// converstion if necessary. This operation is possible by manually placing
// the value of type T in an interface. This syntax would be a convenience.
// ex.  switch v := interface{}(a).(type)

type constraint interface {
	type int, float32, float64
}

func Min[T constraint](a, b T) T {
	switch T {
	case float32:
		return T(math.Min(float64(a), float64(b)))
	case float64:
		return T(math.Min(a, b))
	}
	if a < b {
		return a
	}
	return b
}

// ==============================================================================

func main() {
	a32 := float32(1.3)
	b32 := float32(4.5)
	fmt.Println(Min(a32, b32))
}
