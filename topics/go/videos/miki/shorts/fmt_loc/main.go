// fmt.Printf can access arguments by location.
package main

import (
	"fmt"
)

func main() {
	name, age := "Bugs", 84
	fmt.Printf(
		"%[1]s is %[2]d years old. %[2]d!\n",
		name, age,
	)
	// Bugs is 84 years old. 84!
}
