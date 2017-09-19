package base

// SortDirection specifies sorting direction...
type SortDirection int

const (
	// Descending says that Instances should be sorted high to low...
	Descending SortDirection = 1
	// Ascending states that Instances should be sorted low to high...
	Ascending SortDirection = 2
)

// DataGrid implementations represent data addressable by rows and columns.
type DataGrid interface {
	// Retrieves a given Attribute's specification
	GetAttribute(Attribute) (AttributeSpec, error)
	// Retrieves details of every Attribute
	AllAttributes() []Attribute
	// Marks an Attribute as a class Attribute
	AddClassAttribute(Attribute) error
	// Unmarks an Attribute as a class Attribute
	RemoveClassAttribute(Attribute) error
	// Returns details of all class Attributes
	AllClassAttributes() []Attribute
	// Gets the bytes at a given position or nil
	Get(AttributeSpec, int) []byte
	// Convenience function for iteration.
	MapOverRows([]AttributeSpec, func([][]byte, int) (bool, error)) error
}

// FixedDataGrid implementations have a size known in advance and implement
// all of the functionality offered by DataGrid implementations.
type FixedDataGrid interface {
	DataGrid
	// Returns a string representation of a given row
	RowString(int) string
	// Returns the number of Attributes and rows currently allocated
	Size() (int, int)
}

// UpdatableDataGrid implementations can be changed in addition to implementing
// all of the functionality offered by FixedDataGrid implementations.
type UpdatableDataGrid interface {
	FixedDataGrid
	// Sets a given Attribute and row to a byte sequence.
	Set(AttributeSpec, int, []byte)
	// Adds an Attribute to the grid.
	AddAttribute(Attribute) AttributeSpec
	// Allocates additional room to hold a number of rows
	Extend(int) error
}
