// Use * to dynamically pass flags to fmt verbs.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.*f\n", 2, math.E)
	// 2.72

	fmt.Printf("%.*f\n", 5, math.E)
	// 2.71828

	name := "Fester"
	fmt.Printf("Uncle %-*s!\n", 4, name)
	// Uncle Fester!
	fmt.Printf("Uncle %-*s!\n", 10, name)
	// Uncle Fester    !
}
