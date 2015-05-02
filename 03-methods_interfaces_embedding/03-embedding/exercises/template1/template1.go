// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/N5bME5pkVJ

// Define a nail type with methods drive() and pull(). drive fastens the nail
// into our imaginary board, while pull removes it from the board.
//
// Define a toolbox type that embeds tools, of your choice, that implement
// nailDriver and nailPuller.
package main

type (
	// nailDriver fastens nails
	nailDriver interface {
		driveNail(*nail)
	}

	// nailPuller removes nails
	nailPuller interface {
		pullNail(*nail)
	}
)

// main is the entry point for the application.
func main() {
	// Initialize the toolbox

	// Acquire some nails

	// Display the nails

	// Fasten the nails

	// Now pull half of those nails back out

	// Display the nails again
}
