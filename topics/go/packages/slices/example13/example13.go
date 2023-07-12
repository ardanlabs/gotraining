// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program showcases
// the `slices` package's equal func function.
// The aim of this test is to determine
// if two slices of players are equal.
// This program requires Go 1.21rc1
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Player struct {
	Username string
	Level    int
}

func main() {

	seed := []Player{
		Player{
			Username: "Bill",
			Level:    2,
		},
		Player{
			Username: "Bill",
			Level:    2,
		},
		Player{
			Username: "Alice",
			Level:    2,
		},
	}

	a := slices.Clone(seed)[:2]
	b := slices.Clone(seed)[:2]
	c := slices.Clone(seed)

	// once this function returns false,
	// the two arrays will be deemed
	// different.
	compFunc := func(a, b Player) bool {
		return a.Username == b.Username
	}

	fmt.Println(
		"Is slice a and b equal",
		slices.EqualFunc(a, b, compFunc),
	)

	fmt.Println(
		"Is slice b and c equal",
		slices.EqualFunc(a, c, compFunc),
	)
}
