// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

package game

import "fmt"

// Object represents attributes every object contains.
type Object struct {
	Length float64
	Volume float64
	Mass   float64

	Texture string
	Color   string

	Visible bool
}

// Location represents a position in the game.
type Location struct {
	X int
	Y int
}

// *****************************************************************************

// Drawer is behvior for drawing an object.
type Drawer interface {
	Draw()
}

// Changer is behvior for an object to change shape.
type Changer interface {
	Change(length float64, volume float64, mass float64)
}

// Hider is behavior for an object to hide itself.
type Hider interface {
	Hide(b bool)
}

// Mover is behavior for an object to move.
type Mover interface {
	Move(x int, y int)
}

// *****************************************************************************

// Solid objects contain this behavior.
type Solid interface {
	Drawer
	Hider
	Mover
}

// SolidFixed objects contain this behavior.
type SolidFixed interface {
	Drawer
}

// Liquid objects contain this behavior.
type Liquid interface {
	Drawer
	Changer
	Mover
}

// *****************************************************************************

// DisplaySolid accepts only solid objects and displays
// them in the game.
func DisplaySolid(s Solid) {
	fmt.Println("Solid:", s)
}

// DisplaySolidFixed accepts only solid fixed objects and displays
// them in the game.
func DisplaySolidFixed(sf SolidFixed) {
	fmt.Println("SolidFixed:", sf)
}

// DisplayLiquid accepts only liquid objects and displays
// them in the game.
func DisplayLiquid(l Liquid) {
	fmt.Println("Liquid:", l)
}
