// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/QnkL-UIVJN

// Sample program demonstrating composition through embedding.
package main

import "fmt"

// =============================================================================

// Board represents a surface we can work on.
type Board struct {
	NailsNeeded int
	NailsDriven int
}

// =============================================================================

// NailDriver represents behavior to drive nails into a board.
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

// NailPuller represents behavior to remove nails into a board.
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
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
	b.NailsDriven++

	fmt.Println("Mallet: pounded nail into the board.")
}

// Crowbar is a tool that removes nails.
type Crowbar struct{}

// PullNail yanks a nail out of the specified board.
func (Crowbar) PullNail(nailSupply *int, b *Board) {
	// Yank a nail out of the board.
	b.NailsDriven--

	// Put that nail back into the supply.
	*nailSupply++

	fmt.Println("Crowbar: yanked nail out of the board.")
}

// =============================================================================

// Contractor carries out the task of securing boards.
type Contractor struct{}

// Fasten will drive nails into a board.
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

// Unfasten will remove nails from a board.
func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

// ProcessBoards works against boards.
func (c Contractor) ProcessBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("Contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)

		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(dp, nailSupply, b)
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
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 6},

		// Fresh boards to be fastened.
		{NailsNeeded: 6},
		{NailsNeeded: 9},
		{NailsNeeded: 4},
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
	c.ProcessBoards(&tb, &tb.nails, boards)

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
