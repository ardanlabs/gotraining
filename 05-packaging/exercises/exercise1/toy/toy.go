// Package toy contains support for managing toy inventory.
package toy

// bat represents the bat we sell.
type bat struct {
	Height int
	Weight int
}

// NewBat creates values of type bat.
func NewBat() *bat {
	return new(bat)
}
