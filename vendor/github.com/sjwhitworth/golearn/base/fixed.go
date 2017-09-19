package base

import (
	"bytes"
	"fmt"
)

// FixedAttributeGroups contain a particular number of rows of
// a particular number of Attributes, all of a given type.
type FixedAttributeGroup struct {
	parent     DataGrid
	attributes []Attribute
	size       int
	alloc      []byte
	maxRow     int
}

// String gets a human-readable summary
func (f *FixedAttributeGroup) String() string {
	return "FixedAttributeGroup"
}

// RowSizeInBytes returns the size of each row in bytes
func (f *FixedAttributeGroup) RowSizeInBytes() int {
	return len(f.attributes) * f.size
}

// Attributes returns a slice of Attributes in this FixedAttributeGroup
func (f *FixedAttributeGroup) Attributes() []Attribute {
	ret := make([]Attribute, len(f.attributes))
	// Add Attributes
	for i, a := range f.attributes {
		ret[i] = a
	}
	return ret
}

// AddAttribute adds an attribute to this FixedAttributeGroup
func (f *FixedAttributeGroup) AddAttribute(a Attribute) error {
	f.attributes = append(f.attributes, a)
	return nil
}

// addStorage appends the given storage reference to this FixedAttributeGroup
func (f *FixedAttributeGroup) setStorage(a []byte) {
	f.alloc = a
}

// Storage returns a slice of FixedAttributeGroupStorageRefs which can
// be used to access the memory in this pond.
func (f *FixedAttributeGroup) Storage() []byte {
	return f.alloc
}

func (f *FixedAttributeGroup) offset(col, row int) int {
	return row*f.RowSizeInBytes() + col*f.size
}

func (f *FixedAttributeGroup) set(col int, row int, val []byte) {

	// Double-check the length
	if len(val) != f.size {
		panic(fmt.Sprintf("Tried to call set() with %d bytes, should be %d", len(val), f.size))
	}

	// Find where in the pond the byte is
	offset := f.offset(col, row)

	// Copy the value in
	copied := copy(f.alloc[offset:], val)
	if copied != f.size {
		panic(fmt.Sprintf("set() terminated by only copying %d bytes, should be %d", copied, f.size))
	}

	row++
	if row > f.maxRow {
		f.maxRow = row
	}
}

func (f *FixedAttributeGroup) get(col int, row int) []byte {
	offset := f.offset(col, row)
	return f.alloc[offset : offset+f.size]
}

func (f *FixedAttributeGroup) appendToRowBuf(row int, buffer *bytes.Buffer) {
	for i, a := range f.attributes {
		postfix := " "
		if i == len(f.attributes)-1 {
			postfix = ""
		}
		buffer.WriteString(fmt.Sprintf("%s%s", a.GetStringFromSysVal(f.get(i, row)), postfix))
	}
}

func (f *FixedAttributeGroup) resize(add int) {
	newAlloc := make([]byte, len(f.alloc)+add)
	copy(newAlloc, f.alloc)
	f.alloc = newAlloc
}
