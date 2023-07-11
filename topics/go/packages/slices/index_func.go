// This program showcases
// the `slices` package's index func function.
// The aim of this test is to determine
// the index of the
// specified username.
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
			Level:    2,
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

	indexOf := func(
		name string,
	) func(Player) bool {
		return func(p Player) bool {

			// return true if the name to look for
			// passed is found. this index will be
			// returned to the user.
			if name == p.Username {
				return true
			}
			return false
		}
	}

	indexEron := slices.IndexFunc[[]Player](
		a,
		indexOf("Eron"),
	)

	indexBill := slices.IndexFunc[[]Player](
		a,
		indexOf("Bill"),
	)

	fmt.Println("Eron is at index:", indexEron)
	fmt.Println("Bill is at index:", indexBill)
}
