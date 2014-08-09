// Package toy is for testing
package toy

type bat struct {
	Height int
	Weight int
}

// New Factory function.
func New() *bat {
	return new(bat)
}
