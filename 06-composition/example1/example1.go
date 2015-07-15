// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

package main

import "fmt"

// *****************************************************************************
// Set of behaviors.

// drawer is behvior for drawing an object.
type drawer interface {
	draw()
}

// changer is behvior for an object to change shape.
type changer interface {
	change(length float64, volume float64, mass float64)
}

// hider is behavior for an object to hide itself.
type hider interface {
	hide(b bool)
}

// mover is behavior for an object to move.
type mover interface {
	move(x int, y int)
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
	location
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
	location
	altitude float64
	area     float64
}

// draw is how a cumulus cloud is drawn.
func (c *cumulus) draw() {
	fmt.Printf("[C] %+v\n", *c)
}

// change is how a cumulus cloud can change shape.
func (c *cumulus) change(altitude float64, area float64) {
	c.altitude = altitude
	c.area = area
}

// policeman declares a cop in the game.
type policeman struct {
	location
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
// Set of worlds.

// world represents a set of objects.
type world struct {
	buildings []building
	clouds    []cloud
	people    []person
}

// *****************************************************************************
// Now we can build our world.

func main() {
}
