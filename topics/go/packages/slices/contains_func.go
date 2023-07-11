// This program showcases
// the `slices` package's compact function.
// The aim of this test is to determine
// if a slice contains an element
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

	compareFunc := func(
		arr []Player,
		name string,
	) func(Player) bool {
		return func(p Player) bool {

			// return true if the name to look for
			// passed is found.
			if name == p.Username {
				return true
			}
			return false
		}
	}

	containEron := slices.ContainsFunc[[]Player](
		a,
		compareFunc(a, "Eron"),
	)
	containZack := slices.ContainsFunc[[]Player](
		a,
		compareFunc(a, "Zack"),
	)

	fmt.Println("Does the array contain Eron:", containEron)
	fmt.Println("Does the array contain Zack:", containZack)
}
