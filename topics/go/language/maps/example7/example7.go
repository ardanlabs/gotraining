// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how maps are reference types.
package main

import "fmt"

func main() {

	// Initialize a map with values.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Pass the map to a function to perform some mutation.
	double(scores, "anna")

	// See the change is visible in our map.
	fmt.Println("Score:", scores["anna"])
}

// double finds the score for a specific player and
// multiplies it by 2.
func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
