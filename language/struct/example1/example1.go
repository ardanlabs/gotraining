package main

import (
	"fmt"
)

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare variable of type example and init using
	// a composite literal.
	e := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
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

	fmt.Println(e)
	fmt.Println(anon)
	fmt.Println(anon2)
}
