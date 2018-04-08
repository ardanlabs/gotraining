package base

import (
	"fmt"
)

// AttributeSpec is a pointer to a particular Attribute
// within a particular Instance structure and encodes position
// and storage information associated with that Attribute.
type AttributeSpec struct {
	pond     int
	position int
	attr     Attribute
}

type byPosition []AttributeSpec

func (b byPosition) Len() int {
	return len(b)
}
func (b byPosition) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byPosition) Less(i, j int) bool {

	iPos := (uint64(b[i].pond) << 32) + (uint64(b[i].position))
	jPos := (uint64(b[i].pond) << 32) + (uint64(b[i].position))

	return iPos < jPos
}

// GetAttribute returns an AttributeSpec which matches a given
// Attribute.
func (a *AttributeSpec) GetAttribute() Attribute {
	return a.attr
}

// String returns a human-readable description of this AttributeSpec.
func (a *AttributeSpec) String() string {
	return fmt.Sprintf("AttributeSpec(Attribute: '%s', Pond: %d/%d)", a.attr, a.pond, a.position)
}
