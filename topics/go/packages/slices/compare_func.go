// This program showcases
// the `slices` package's compare func function.
// The aim of this test is to leverage
// the compare function to determine
// which array's length is greater
// This program requires Go 1.21rc1
package main

import (
	"fmt"
	"slices"
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
			Username: "Alice",
			Level:    2,
		},
		Player{
			Username: "Zack",
			Level:    4,
		},
		Player{
			Username: "Eron",
			Level:    3,
		},
	}
	a := slices.Clone(seed)[:3]

	b := slices.Clone(seed)[:2]

	c := slices.Clone(seed)

	d := slices.Clone(seed)

	comp := func(a, b Player) int {
		if a.Username == "" {
			return -1
		}

		if b.Username == "" {
			return 1
		}

		return 0
	}

	// d is short for
	// dictionary and translates
	// the output from the compare
	// function into something that is
	// human readable.
	dict := map[int]string{
		-1: "First slice is shorter",
		0:  "Both slices are equal",
		1:  "Second slice is shorter",
	}
	fmt.Println(
		"Compare Slice a and b",
		dict[slices.CompareFunc(a, b, comp)],
	)

	fmt.Println(
		"Compare Slice a and c",
		dict[slices.CompareFunc(a, c, comp)],
	)

	fmt.Println(
		"Compare Slice c and d",
		dict[slices.CompareFunc(c, d, comp)],
	)
}
