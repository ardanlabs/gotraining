package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/06-composition/example1/game"
)

type Building struct {
	name string
	game.Object
	game.Location
}

func (b *Building) Draw() {
	fmt.Printf("[B] %+v\n", *b)
}

type Cloud struct {
	kind string
	game.Object
	game.Location
}

func (c *Cloud) Draw() {
	fmt.Printf("[C] %+v\n", *c)
}

func (c *Cloud) Change(length float64, volume float64, mass float64) {
	c.Length = length
	c.Volume = volume
	c.Mass = mass
}

func (c *Cloud) Move(x int, y int) {
	c.X += x
	c.Y += y
}

type Player struct {
	name string
	game.Object
	game.Location
}

func (p *Player) Draw() {
	fmt.Printf("[P] %+v\n", *p)
}

func (p *Player) Move(x int, y int) {
	p.X = x
	p.Y = y
}

func (p *Player) Hide(b bool) {
	p.Visible = !b
}

func main() {
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

	game.Draw(&b)

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

	game.Draw(&c)
	game.Move(&c, 30, 12)

	p := Player{
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

	game.Draw(&p)
	game.Move(&p, 300, 123)
	game.Hide(&p, false)
}
