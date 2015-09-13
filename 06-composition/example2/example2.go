// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/we7a-Nk08J

// Sample program demonstrating composition through embedding.
package main

import "fmt"

// =============================================================================

// Board represents a surface we can work on.
type Board struct {
	nailsNeeded int
	nailsDriven int
}

// =============================================================================

// NailDriver represents behavior to drive nails into a board.
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

// NailPuller represents behavior to remove nails into a board.
type NailPuller interface {
	pullNail(nailSupply *int, b *Board)
}

// NailDrivePuller represents behavior to drive and remove nails into a board.
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// =============================================================================

// Mallet is a tool that pounds in nails.
type Mallet struct{}

// DriveNail pounds a nail into the specified board.
func (Mallet) DriveNail(nailSupply *int, b *Board) {
	// Take a nail out of the supply.
	*nailSupply--

	// Pound a nail into the board.
	b.nailsDriven++

	fmt.Println("Mallet: pounded nail into the board.")
}

// Crowbar is a tool that removes nails.
type Crowbar struct{}

// pullNail yanks a nail out of the specified board.
func (Crowbar) pullNail(nailSupply *int, b *Board) {
	// Yank a nail out of the board.
	b.nailsDriven--

	// Put that nail back into the supply.
	*nailSupply++

	fmt.Println("Crowbar: yanked nail out of the board.")
}

// =============================================================================

// Contractor carries out the task of securing boards.
type Contractor struct{}

// fasten will drive nails into a board.
func (Contractor) fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.nailsDriven < b.nailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

// unfasten will remove nails from a board.
func (Contractor) unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.nailsDriven > b.nailsNeeded {
		p.pullNail(nailSupply, b)
	}
}

// processBoards works against boards.
func (c Contractor) processBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("Contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.nailsDriven < b.nailsNeeded:
			c.fasten(dp, nailSupply, b)

		case b.nailsDriven > b.nailsNeeded:
			c.unfasten(dp, nailSupply, b)
		}
	}
}

// =============================================================================

// Toolbox can contains any number of tools.
type Toolbox struct {
	NailDriver
	NailPuller

	nails int
}

// =============================================================================

// main is the entry point for the application.
func main() {
	// Inventory of old boards to remove, and the new boards
	// that will replace them.
	boards := []Board{
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
	tb := Toolbox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails:      10,
	}

	// Display the current state of our toolbox and boards.
	displayState(&tb, boards)

	// Hire a Contractor and put our Contractor to work.
	var c Contractor
	c.processBoards(&tb, &tb.nails, boards)

	// Display the new state of our toolbox and boards.
	displayState(&tb, boards)
}

// displayState provide information about all the boards.
func displayState(tb *Toolbox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}
