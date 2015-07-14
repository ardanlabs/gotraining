// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/QKIjbBDW16

// Sample program demonstrating composition through embedding.
package main

import "fmt"

// board represents a surface we can work on.
type board struct {
	nailsNeeded int
	nailsDriven int
}

// A set of interfaces for driving and pulling nails.
type (
	nailDriver interface {
		driveNail(nailSupply *int, b *board)
	}

	nailPuller interface {
		pullNail(nailSupply *int, b *board)
	}

	nailDrivePuller interface {
		nailDriver
		nailPuller
	}
)

// mallet is a tool that pounds in nails.
type mallet struct{}

// driveNail pounds a nail into the specified board.
func (mallet) driveNail(nailSupply *int, b *board) {
	// Take a nail out of the supply.
	*nailSupply--

	// Pound a nail into the board.
	b.nailsDriven++

	fmt.Println("mallet: pounded nail into the board.")
}

// crowbar is a tool that removes nails.
type crowbar struct{}

// pullNail yanks a nail out of the specified board.
func (crowbar) pullNail(nailSupply *int, b *board) {
	// Yank a nail out of the board.
	b.nailsDriven--

	// Put that nail back into the supply.
	*nailSupply++

	fmt.Println("crowbar: yanked nail out of the board.")
}

// toolbox can contains any number of tools.
type toolbox struct {
	nailDriver
	nailPuller

	nails int
}

// contractor carries out the task of securing boards.
type contractor struct{}

// fasten will drive nails into a board.
func (contractor) fasten(d nailDriver, nailSupply *int, b *board) {
	for b.nailsDriven < b.nailsNeeded {
		d.driveNail(nailSupply, b)
	}
}

// unfasten will remove nails from a board.
func (contractor) unfasten(p nailPuller, nailSupply *int, b *board) {
	for b.nailsDriven > b.nailsNeeded {
		p.pullNail(nailSupply, b)
	}
}

// processBoards works against boards.
func (c contractor) processBoards(dp nailDrivePuller, nailSupply *int, boards []board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.nailsDriven < b.nailsNeeded:
			c.fasten(dp, nailSupply, b)

		case b.nailsDriven > b.nailsNeeded:
			c.unfasten(dp, nailSupply, b)
		}
	}
}

// displayState provide information about all the boards.
func displayState(tb *toolbox, boards []board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}
}

// main is the entry point for the application.
func main() {
	// Inventory of old boards to remove, and the new boards
	// that will replace them.
	boards := []board{
		// Rotted boards to be removed.
		{nailsDriven: 3},
		{nailsDriven: 1},
		{nailsDriven: 6},

		// Fresh boards to be fastened.
		{nailsNeeded: 6},
		{nailsNeeded: 9},
		{nailsNeeded: 4},
	}

	// Fill a toolbox.
	tb := toolbox{
		nailDriver: mallet{},
		nailPuller: crowbar{},
		nails:      10,
	}

	// Display the current state of our toolbox and boards.
	displayState(&tb, boards)

	fmt.Println()

	// Hire a contractor and put our contractor to work.
	var c contractor
	c.processBoards(&tb, &tb.nails, boards)

	fmt.Println()

	// Display the new state of our toolbox and boards.
	displayState(&tb, boards)
}
