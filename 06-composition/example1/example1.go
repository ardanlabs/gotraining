// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/W5ya6_LAU6

// Sample program demonstrating composition through embedding.
package main

import "fmt"

// *****************************************************************************
// Set of behaviors.

// drawer is behavior for drawing an object.
type drawer interface {
	draw()
}

// changer is behavior for an object to change shape.
type changer interface {
	change(area float64)
}

// hider is behavior for an object to hide itself.
type hider interface {
	hide(b bool)
}

// mover is behavior for an object to move.
type mover interface {
	move(x, y, z int)
}

// *****************************************************************************
// Set of objects that exhibit different behaviors.

// building contain this behavior.
type building interface {
	drawer
}

// cloud contain this behavior.
type cloud interface {
	drawer
	changer
	mover
}

// person contain this behavior.
type person interface {
	drawer
	hider
	mover
}

// *****************************************************************************
// Set of concrete objects.

// location provides support placing objects.
type location struct {
	x, y, z int
}

// move is how a cumulus cloud can move.
func (l *location) move(x, y, z int) {
	l.x = x
	l.y = y
	l.z = z
}

// house represents a place people live in.
type house struct {
	*location
	address string
	color   string
}

// draw is how a house is drawn.
func (h *house) draw() {
	fmt.Printf("[H] %+v\n", *h)
}

// move is implemented to prevent a house from moving.
func (*house) move(x, y, z int) {}

// cumulus declares a cumulus cloud in the game.
type cumulus struct {
	*location
	area float64
}

// draw is how a cumulus cloud is drawn.
func (c *cumulus) draw() {
	fmt.Printf("[C] %+v\n", *c)
}

// change is how a cumulus cloud can change shape.
func (c *cumulus) change(area float64) {
	c.area = area
}

// policeman declares a cop in the game.
type policeman struct {
	*location
	name    string
	visible bool
}

// draw is how a cop is drawn.
func (p *policeman) draw() {
	fmt.Printf("[P] %+v\n", *p)
}

// hide is how a cop can become hidden.
func (p *policeman) hide(b bool) {
	p.visible = !b
}

// *****************************************************************************
// Declare a world that contains all these behaviors.

// world represents a set of objects.
type world struct {
	buildings []building
	clouds    []cloud
	people    []person
}

// *****************************************************************************
// Now we can build our world.

// main is the entry point for the application.
func main() {
	// Create a world.
	w := world{
		buildings: []building{
			&house{&location{10, 10, 10}, "123 mocking bird", "white"},
			&house{&location{10, 20, 10}, "127 mocking bird", "red"},
		},
		clouds: []cloud{
			&cumulus{&location{10, 15, 1000}, 123456.4332},
		},
		people: []person{
			&policeman{&location{15, 40, 10}, "Harry", true},
		},
	}

	// Build a slice of all the values that can be drawn.
	var d []drawer
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
func draw(d []drawer) {
	for _, v := range d {
		v.draw()
	}
}
