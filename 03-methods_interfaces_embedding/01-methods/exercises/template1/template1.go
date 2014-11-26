// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/uYkS_J8LuL

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// batter represents a playing in the game.
/*
	type type_name struct {
		field_name type
		field_name type
		field_name type
	}
*/

// main is the entry point for the application.
func main() {
	// Create a few players.
	/*
		slice_name := []type_name{
			type_name{value, value, value},
			type_name{value, value, value},
			type_name{value, value, value},
		}
	*/

	// Display the batting average for each player.
	/*
		for _, variable_name := range slice_name {
			fmt.Printf("%s: AVG[%.3f]\n", variable_name.field_name, variable_name.method_name())
		}
	*/
}

// average calculates the batting average for a batter.
/*
	func (receiver_name *type_name) function_name() return_type {
		return float64(receiver_name.field_name) / float64(receiver_name.field_name)
	}
*/
