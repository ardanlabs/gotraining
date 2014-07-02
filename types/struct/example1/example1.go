package main

import (
	"fmt"
)

type Example struct {
	BoolValue  bool
	IntValue   int16
	FloatValue float32
}

func main() {
	// Declare variable of type Example and init using
	// a composite literal.
	example := Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
	}

	// Declare a variable of an anonymous type and init
	// using a composite literal.
	anon := struct {
		name string
	}{
		name: "Jill",
	}

	// Declare a variable of an anonymous type that contains
	// an anonymous inner type and init using a composite literal.
	anon2 := struct {
		inner struct {
			name string
		}
		age int
	}{
		inner: struct {
			name string
		}{"Bill"},
		age: 45,
	}

	fmt.Println(example)
	fmt.Println(anon)
	fmt.Println(anon2)
}
