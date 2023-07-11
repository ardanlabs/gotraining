// This program showcases
// the `slices` package's is sorted func function
// to determine if an array is in ascending
// order.
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

	a := []Player{
		Player{
			Username: "Bill",
			Level:    6,
		},
		Player{
			Username: "Alice",
			Level:    2,
		},
		Player{
			Username: "Eron",
			Level:    3,
		},
	}

	b := []Player{
		Player{
			Username: "Bill",
			Level:    1,
		},
		Player{
			Username: "Alice",
			Level:    2,
		},
		Player{
			Username: "Eron",
			Level:    3,
		},
	}

	// if a's level is greater than b,
	// the array is not sorted in ascending order/
	cmp := func(a, b Player) int {
		return a.Level - b.Level
	}

	isSortedA := slices.IsSortedFunc(a, cmp)
	isSortedB := slices.IsSortedFunc(b, cmp)

	fmt.Println("Is Array A in ascending order:", isSortedA)
	fmt.Println("Is Array B in ascending order:", isSortedB)
}
