// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// player represents a person in the game.
type player struct {
	name   string
	atBats int
	hits   int
}

// average calculates the batting average for a player.
func (p *player) average() float64 {
	if p.atBats == 0 {
		return 0.0
	}

	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a few players.
	ps := []player{
		{"bill", 10, 7},
		{"jim", 12, 6},
		{"ed", 6, 4},
	}

	// Display the batting average for each player.
	for i := range ps {
		fmt.Printf("%s: AVG[.%.f]\n", ps[i].name, ps[i].average()*1000)
	}

	// Why did I not choose this form?
	for _, p := range ps {
		fmt.Printf("%s: AVG[.%.f]\n", p.name, p.average()*1000)
	}
}
