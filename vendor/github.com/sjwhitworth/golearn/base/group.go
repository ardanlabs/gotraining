package base

import (
	"bytes"
)

// AttributeGroups store related sequences of system values
// in memory for the DenseInstances structure.
type AttributeGroup interface {
	// Used for printing
	appendToRowBuf(row int, buffer *bytes.Buffer)
	// Adds a new Attribute
	AddAttribute(Attribute) error
	// Returns all Attributes
	Attributes() []Attribute
	// Gets the byte slice at a given column, row offset
	get(int, int) []byte
	// Stores the byte slice at a given column, row offset
	set(int, int, []byte)
	// Sets the reference to underlying memory
	setStorage([]byte)
	// Gets the size of each row in bytes (rounded up)
	RowSizeInBytes() int
	// Adds some storage to this group
	resize(int)
	// Gets a reference to underlying memory
	Storage() []byte
	// Returns a human-readable summary
	String() string
}
