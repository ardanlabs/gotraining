// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

package game

import "fmt"

type Object struct {
	Length float64
	Volume float64
	Mass   float64

	Texture string
	Color   string

	Visible bool
}

type Location struct {
	X int
	Y int
}

// *****************************************************************************

type Drawer interface {
	Draw()
}

type Changer interface {
	Change(length float64, volume float64, mass float64)
}

type Hider interface {
	Hide(b bool)
}

type Mover interface {
	Move(x int, y int)
}

// *****************************************************************************

type Solid interface {
	Drawer
	Hider
	Mover
}

type SolidFixed interface {
	Drawer
}

type Liquid interface {
	Drawer
	Changer
	Mover
}

// *****************************************************************************

func DisplaySolid(s Solid) {
	fmt.Println("Solid:", s)
}

func DisplaySolidFixed(sf SolidFixed) {
	fmt.Println("SolidFixed:", sf)
}

func DisplayLiquid(l Liquid) {
	fmt.Println("Liquid:", l)
}
