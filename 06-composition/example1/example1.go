// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/DqZvyTbTle

// Sample program demonstrating composition through embedding.
package main

import "fmt"

// *****************************************************************************
// Set of behaviors.

// Drawer is behavior for drawing an object.
type Drawer interface {
	Draw()
}

// Changer is behavior for an object to change shape.
type Changer interface {
	Change(area float64)
}

// Hider is behavior for an object to hide itself.
type Hider interface {
	Hide(b bool)
}

// Mover is behavior for an object to move.
type Mover interface {
	Move(x, y, z int)
}

// *****************************************************************************
// Set of objects that exhibit different behaviors.

// Building contain this behavior.
type Building interface {
	Drawer
}

// Cloud contain this behavior.
type Cloud interface {
	Drawer
	Changer
	Mover
}

// Person contain this behavior.
type Person interface {
	Drawer
	Hider
	Mover
}

// *****************************************************************************
// Set of concrete objects.

// Location provides support placing objects.
type Location struct {
	X, Y, Z int
}

// Move is how a cumulus cloud can move.
func (l *Location) Move(x, y, z int) {
	l.X = x
	l.Y = y
	l.Z = z
}

// House represents a place people live in.
type House struct {
	Location
	Address string
	Color   string
}

// Draw is how a House is drawn.
func (h *House) Draw() {
	fmt.Printf("[H] %+v\n", *h)
}

// Move is implemented to prevent a House from moving.
func (*House) Move(x, y, z int) {}

// Cumulus declares a cumulus cloud in the game.
type Cumulus struct {
	Location
	Area float64
}

// Draw is how a cumulus cloud is drawn.
func (c *Cumulus) Draw() {
	fmt.Printf("[C] %+v\n", *c)
}

// Change is how a cumulus cloud can change shape.
func (c *Cumulus) Change(area float64) {
	c.Area = area
}

// Policeman declares a cop in the game.
type Policeman struct {
	Location
	Name    string
	Visible bool
}

// Draw is how a cop is drawn.
func (p *Policeman) Draw() {
	fmt.Printf("[P] %+v\n", *p)
}

// Hide is how a cop can become hidden.
func (p *Policeman) Hide(b bool) {
	p.Visible = !b
}

// *****************************************************************************
// Declare a world that contains all these behaviors.

// world represents a set of objects.
type world struct {
	buildings []Building
	clouds    []Cloud
	people    []Person
}

// *****************************************************************************
// Now we can build our world.

// main is the entry point for the application.
func main() {
	// Create a world.
	w := world{
		buildings: []Building{
			&House{Location{10, 10, 10}, "123 mocking bird", "white"},
			&House{Location{10, 20, 10}, "127 mocking bird", "red"},
		},
		clouds: []Cloud{
			&Cumulus{Location{10, 15, 1000}, 123456.4332},
		},
		people: []Person{
			&Policeman{Location{15, 40, 10}, "Harry", true},
		},
	}

	// Build a slice of all the values that can be drawn.
	var d []Drawer
	for _, v := range w.buildings {
		d = append(d, v)
	}
	for _, v := range w.clouds {
		d = append(d, v)
	}
	for _, v := range w.people {
		d = append(d, v)
	}

	// Draw the world.
	draw(d)
}

// draw takes a set of values that can be drawn.
func draw(d []Drawer) {
	for _, v := range d {
		v.Draw()
	}
}
