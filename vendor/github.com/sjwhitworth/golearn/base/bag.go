package base

import (
	"bytes"
	"fmt"
)

// BinaryAttributeGroups contain only BinaryAttributes
// Compact each Attribute to a bit for better storage
type BinaryAttributeGroup struct {
	parent     DataGrid
	attributes []Attribute
	size       int
	alloc      []byte
	maxRow     int
}

// String returns a human-readable summary.
func (b *BinaryAttributeGroup) String() string {
	return "BinaryAttributeGroup"
}

// RowSizeInBytes returns the size of each row in bytes
// (rounded up to nearest byte).
func (b *BinaryAttributeGroup) RowSizeInBytes() int {
	return (len(b.attributes) + 7) / 8
}

// Attributes returns a slice of Attributes in this BinaryAttributeGroup.
func (b *BinaryAttributeGroup) Attributes() []Attribute {
	ret := make([]Attribute, len(b.attributes))
	for i, a := range b.attributes {
		ret[i] = a
	}
	return ret
}

// AddAttribute adds an Attribute to this BinaryAttributeGroup
func (b *BinaryAttributeGroup) AddAttribute(a Attribute) error {
	b.attributes = append(b.attributes, a)
	return nil
}

// Storage returns a reference to the underlying storage.
//
// IMPORTANT: don't modify
func (b *BinaryAttributeGroup) Storage() []byte {
	return b.alloc
}

//
// internal methods
//

func (b *BinaryAttributeGroup) setStorage(a []byte) {
	b.alloc = a
}

func (b *BinaryAttributeGroup) getByteOffset(col, row int) int {
	return row*b.RowSizeInBytes() + col/8
}

func (b *BinaryAttributeGroup) set(col, row int, val []byte) {

	offset := b.getByteOffset(col, row)

	// If the value is 1, OR it
	if val[0] > 0 {
		b.alloc[offset] |= (1 << (uint(col) % 8))
	} else {
		// Otherwise, AND its complement
		b.alloc[offset] &= ^(1 << (uint(col) % 8))
	}

	row++
	if row > b.maxRow {
		b.maxRow = row
	}
}

func (b *BinaryAttributeGroup) get(col, row int) []byte {
	offset := b.getByteOffset(col, row)
	if b.alloc[offset]&(1<<(uint(col%8))) > 0 {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func (b *BinaryAttributeGroup) appendToRowBuf(row int, buffer *bytes.Buffer) {
	for i, a := range b.attributes {
		postfix := " "
		if i == len(b.attributes)-1 {
			postfix = ""
		}
		buffer.WriteString(fmt.Sprintf("%s%s",
			a.GetStringFromSysVal(b.get(i, row)), postfix))
	}
}

func (b *BinaryAttributeGroup) resize(add int) {
	newAlloc := make([]byte, len(b.alloc)+add)
	copy(newAlloc, b.alloc)
	b.alloc = newAlloc
}
