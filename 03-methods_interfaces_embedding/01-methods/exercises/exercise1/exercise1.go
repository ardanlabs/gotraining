// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/C8Z_MiYKbc

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// batter represents a playing in the game.
type batter struct {
	name   string
	atBats int
	hits   int
}

// average calculates the batting average for a batter.
func (b *batter) average() float64 {
	return float64(b.hits) / float64(b.atBats)
}

// main is the entry point for the application.
func main() {
	// Create a few players.
	players := []batter{
		{"bill", 10, 7},
		{"jim", 12, 6},
		{"ed", 6, 4},
	}

	// Display the batting average for each player.
	for _, player := range players {
		average := player.average() * 1000
		fmt.Printf("%s: AVG[.%.f]\n", player.name, average)
	}
}
