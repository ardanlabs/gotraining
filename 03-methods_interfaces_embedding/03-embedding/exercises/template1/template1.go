// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/N5bME5pkVJ

// Define a toolbox type that embeds tools, of your choice, that implement
// nailDriver and nailPuller.
package main

type nail struct {
	// declare a field to manage nail state (fastened or unfastened)
	// nails are driven into our implicit, imaginary board
}

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

type toolbox struct {
	// _embed_ your tool(s) here
}

// main is the entry point for the application.
func main() {
	// Initialize the toolbox

	// Acquire some nails

	// Print the slice of nails

	// Fasten the nails (pass in the toolbox, not the tool)

	// Now pull half of those nails back out

	// Print the slice of nails again (pass in the toolbox, not the tool)
}

func buyNails(n int) []*nail {
	// return a slice of n new nails
}

func driveNails(d nailDriver, nails []*nail) {
	// drive the specified nails using the provided nailDriver
}

func pullNails(p nailPuller, nails []*nail) {
	// pull the specified nails using the provided nailPuller
}
