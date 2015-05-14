// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/TW2E6ApbTD

// Define a nail type with methods drive() and pull(). drive fastens the nail
// into our imaginary board, while pull removes it from the board.
//
// Define a toolbox type that embeds tools, of your choice, that implement
// nailDriver and nailPuller.
package main

import "fmt"

type nail struct {
	// fastened is true if pinned, false otherwise
	fastened bool
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

func (n *nail) drive() { n.fastened = true }
func (n *nail) pull()  { n.fastened = false }

func (n *nail) String() string {
	s := "-"
	if n.fastened {
		s = "+"
	}
	return fmt.Sprintf("%s%p", s, n)
}

type clawhammer struct{}

func (h clawhammer) driveNail(n *nail) { n.drive() }
func (h clawhammer) pullNail(n *nail)  { n.pull() }

type toolbox struct {
	clawhammer
}

// main is the entry point for the application.
func main() {
	// This toolbox comes with a free hammer
	var b toolbox

	const n = 4

	// Acquire some nails
	nails := buyNails(n)

	// Display the nails
	fmt.Println(nails)

	// Fasten the nails
	driveNails(b, nails)

	// Now pull half of those nails back out
	pullNails(b, nails[n/2:])

	// Display the nails again
	fmt.Println(nails)
}

func buyNails(n int) []*nail {
	nails := make([]*nail, n)
	for i := range nails {
		nails[i] = new(nail)
	}
	return nails
}

func driveNails(d nailDriver, nails []*nail) {
	for _, n := range nails {
		d.driveNail(n)
	}
}

func pullNails(p nailPuller, nails []*nail) {
	for _, n := range nails {
		p.pullNail(n)
	}
}
