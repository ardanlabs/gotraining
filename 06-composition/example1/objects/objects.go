// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package objects contains game related objects and behavior.
package objects

import "fmt"

// Building represents a building in the game.
type Building struct {
	location
	solid
}

// NewBuilding creates a new building and places it in the game.
func NewBuilding(x, y int, shape, color string) *Building {
	var b Building
	b.Set(x, y)
	b.Draw(shape, color)
	return &b
}

// Paint provides a visual of the building and its location.
func (b *Building) Paint() {
	fmt.Printf("[B]%s%s\n", b.location.Paint(), b.solid.Paint())
}

// Cloud represents a cloud in the game.
type Cloud struct {
	movable
}

// NewCloud creates a new cloud and places it in the game.
func NewCloud(x, y int) *Cloud {
	var c Cloud
	c.Set(x, y)
	return &c
}

// Paint provides a visual of the cloud and its location.
func (c *Cloud) Paint() {
	fmt.Printf("[C]%s\n", c.location.Paint())
}

// Player represents a player in the game.
type Player struct {
	movable
	solid
}

// NewPlayer creates a new player and places it in the game.
func NewPlayer(x, y int, shape, color string) *Player {
	var p Player
	p.Set(x, y)
	p.Draw(shape, color)
	return &p
}

// Paint provides a visual of the player and its location.
func (p *Player) Paint() {
	fmt.Printf("[P]%s%s\n", p.location.Paint(), p.solid.Paint())
}
