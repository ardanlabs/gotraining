package game

import (
	"fmt"
)

// Key is a key in the game
type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

// implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

// Player is a player in the game
type Player struct {
	Name string
	Keys []Key
}
