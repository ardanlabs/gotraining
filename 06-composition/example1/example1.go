// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program demonstrating composition through embedding.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/06-composition/example1/game"
)

// Building declares a building in the game.
type Building struct {
	name string
	game.Object
	game.Location
}

// Draw is how a building is drawn.
func (b *Building) Draw() {
	fmt.Printf("[B] %+v\n", *b)
}

// *****************************************************************************

// Cloud declares a cloud in the game.
type Cloud struct {
	kind string
	game.Object
	game.Location
}

// Draw is how a cloud is drawn.
func (c *Cloud) Draw() {
	fmt.Printf("[C] %+v\n", *c)
}

// Change is how a cloud can change shape.
func (c *Cloud) Change(length float64, volume float64, mass float64) {
	c.Length = length
	c.Volume = volume
	c.Mass = mass
}

// Move is how a cloud can move.
func (c *Cloud) Move(x int, y int) {
	c.X += x
	c.Y += y
}

// *****************************************************************************

// Person declares a person in the game.
type Person struct {
	name string
	game.Object
	game.Location
}

// Draw is how a person is drawn.
func (p *Person) Draw() {
	fmt.Printf("[P] %+v\n", *p)
}

// Move is how a person moves.
func (p *Person) Move(x int, y int) {
	p.X = x
	p.Y = y
}

// Hide is how a person can become invisible.
func (p *Person) Hide(b bool) {
	p.Visible = !b
}

// *****************************************************************************

// main is the entry point for the application.
func main() {
	// Create a building.
	b := Building{
		name: "NY Times",
		Object: game.Object{
			Length:  100000,
			Volume:  37e6,
			Mass:    85e9,
			Texture: "stone",
			Color:   "silver",
		},
		Location: game.Location{
			X: 80,
			Y: 64,
		},
	}

	// Create a cloud.
	c := Cloud{
		kind: "cirrus",
		Object: game.Object{
			Length:  5000,
			Volume:  4e10,
			Mass:    8818490,
			Texture: "puffy",
			Color:   "white",
		},
		Location: game.Location{
			X: 13280,
			Y: 33464,
		},
	}

	// Create a person.
	p := Person{
		name: "Bill",
		Object: game.Object{
			Length:  72,
			Volume:  66.4,
			Mass:    68.0,
			Texture: "skin",
			Color:   "white",
		},
		Location: game.Location{
			X: 13280,
			Y: 33464,
		},
	}

	// Using the game display functions, pass the
	// correct object type through. The compiler will
	// provide checks based on interface implementation.

	game.DisplaySolidFixed(&b)
	game.DisplayLiquid(&c)
	game.DisplaySolid(&p)
}
