package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// =============================================================================

// This code defines two user defined types based on an underlying concrete type.
// Each type implements a method named last that returns the value stored at the
// highest index position in the vector or an error when the vector is empty.

type vectorInt []int

func (v vectorInt) last() (int, error) {
	if len(v) == 0 {
		return 0, errors.New("empty")
	}
	return v[len(v)-1], nil
}

type vectorString []string

func (v vectorString) last() (string, error) {
	if len(v) == 0 {
		return "", errors.New("empty")
	}
	return v[len(v)-1], nil
}

// =============================================================================

// This user defined type is based on the empty interface so any value can be
// added into the vector. Since the last function is using the empty interface
// for the return, users will need to perform type assertions to get
// back to the concrete value being stored inside the interface.

type vectorInterface []interface{}

func (v vectorInterface) last() (interface{}, error) {
	if len(v) == 0 {
		return nil, errors.New("empty")
	}
	return v[len(v)-1], nil
}

// =============================================================================

// This is a generics version of the user defined type. It represents a slice of
// some type T (to be determined later). The last method is declared with a 
// value receiver based on a vector of the same type T and returns a value
// of that same type T as well.

type vector[T any] []T

func (v vector[T]) last() (T, error) {
	var zero T
	if len(v) == 0 {
		return zero, errors.New("empty")
	}
	return v[len(v)-1], nil
}

// =============================================================================

func main() {
	fmt.Print("vectorInt : ")
	vInt := vectorInt{10, -1}
	i, err := vInt.last()
	if i < 0 {
		fmt.Print("negative integer: ")
	}
	fmt.Printf("value: %d error: %v\n", i, err)

	fmt.Print("vectorString : ")
	vStr := vectorString{"A", "B", string([]byte{0xff})}
	s, err := vStr.last()
	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}
	fmt.Printf("value: %q error: %v\n", s, err)

	// =========================================================================

	fmt.Print("vectorInterface : ")
	vItf := vectorInterface{10, "A", 20, "B", 3.14}
	itf, err := vItf.last()
	switch v := itf.(type) {
	case int:
		if v < 0 {
			fmt.Print("negative integer: ")
		}
	case string:
		if !utf8.ValidString(v) {
			fmt.Print("non-valid string: ")
		}
	default:
		fmt.Printf("unknown type %T: ", v)
	}
	fmt.Printf("value: %v error: %v\n", itf, err)

	// =========================================================================

	fmt.Print("vector[int] : ")
	vGenInt := vector[int]{10, -1}
	i, err = vGenInt.last()
	if i < 0 {
		fmt.Print("negative integer: ")
	}
	fmt.Printf("value: %d error: %v\n", i, err)

	fmt.Print("vector[string] : ")
	vGenStr := vector[string]{"A", "B", string([]byte{0xff})}
	s, err = vGenStr.last()
	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}
	fmt.Printf("value: %q error: %v\n", s, err)
}
