// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how maps behave when you read an
// absent key.
package main

import "fmt"

func main() {

	// Create a map to track scores for players in a game.
	scores := make(map[string]int)

	// Read the element at key "anna". It is absent so we get
	// the zero-value for this map's value type.
	score := scores["anna"]

	fmt.Println("Score:", score)

	// If we need to check for the presence of a key we use
	// a 2 variable assignment. The 2nd variable is a bool.
	score, ok := scores["anna"]

	fmt.Println("Score:", score, "Present:", ok)

	// We can leverage the zero-value behavior to write
	// convenient code like this:
	scores["anna"]++

	// Without this behavior we would have to code in a
	// defensive way like this:
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)
}
