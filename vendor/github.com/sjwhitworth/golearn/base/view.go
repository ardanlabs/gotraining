package base

import (
	"bytes"
	"fmt"
)

// InstancesViews hide or re-order Attributes and rows from
// a given DataGrid to make it appear that they've been deleted.
type InstancesView struct {
	src        FixedDataGrid
	attrs      []AttributeSpec
	rows       map[int]int
	classAttrs map[Attribute]bool
	maskRows   bool
}

func (v *InstancesView) addClassAttrsFromSrc(src FixedDataGrid) {
	for _, a := range src.AllClassAttributes() {
		matched := true
		if v.attrs != nil {
			matched = false
			for _, b := range v.attrs {
				if b.attr.Equals(a) {
					matched = true
				}
			}
		}
		if matched {
			v.classAttrs[a] = true
		}
	}
}

func (v *InstancesView) resolveRow(origRow int) int {
	if v.rows != nil {
		if newRow, ok := v.rows[origRow]; !ok {
			if v.maskRows {
				return -1
			}
		} else {
			return newRow
		}
	}

	return origRow

}

// NewInstancesViewFromRows creates a new InstancesView from a source
// FixedDataGrid and row -> row mapping. The key of the rows map is the
// row as it exists within this mapping: for example an entry like 5 -> 1
// means that row 1 in src will appear at row 5 in the Instancesview.
//
// Rows are not masked in this implementation, meaning that all rows which
// are left unspecified appear as normal.
func NewInstancesViewFromRows(src FixedDataGrid, rows map[int]int) *InstancesView {
	ret := &InstancesView{
		src,
		nil,
		rows,
		make(map[Attribute]bool),
		false,
	}

	ret.addClassAttrsFromSrc(src)
	return ret
}

// NewInstancesViewFromVisible creates a new InstancesView from a source
// FixedDataGrid, a slice of row numbers and a slice of Attributes.
//
// Only the rows specified will appear in this InstancesView, and they will
// appear in the same order they appear within the rows array.
//
// Only the Attributes specified will appear in this InstancesView. Retrieving
// Attribute specifications from this InstancesView will maintain their order.
func NewInstancesViewFromVisible(src FixedDataGrid, rows []int, attrs []Attribute) *InstancesView {
	ret := &InstancesView{
		src,
		ResolveAttributes(src, attrs),
		make(map[int]int),
		make(map[Attribute]bool),
		true,
	}

	for i, a := range rows {
		ret.rows[i] = a
	}

	ret.addClassAttrsFromSrc(src)
	return ret
}

// NewInstancesViewFromAttrs creates a new InstancesView from a source
// FixedDataGrid and a slice of Attributes.
//
// Only the Attributes specified will appear in this InstancesView.
func NewInstancesViewFromAttrs(src FixedDataGrid, attrs []Attribute) *InstancesView {
	ret := &InstancesView{
		src,
		ResolveAttributes(src, attrs),
		nil,
		make(map[Attribute]bool),
		false,
	}

	ret.addClassAttrsFromSrc(src)
	return ret
}

// GetAttribute returns an Attribute specification matching an Attribute
// if it has not been filtered.
//
// The AttributeSpecs returned are the same as those returned by the
// source FixedDataGrid.
func (v *InstancesView) GetAttribute(a Attribute) (AttributeSpec, error) {
	if a == nil {
		return AttributeSpec{}, fmt.Errorf("Attribute can't be nil")
	}
	// Pass-through on nil
	if v.attrs == nil {
		return v.src.GetAttribute(a)
	}
	// Otherwise
	for _, r := range v.attrs {
		// If the attribute matches...
		if r.GetAttribute().Equals(a) {
			return r, nil
		}
	}
	return AttributeSpec{}, fmt.Errorf("Requested Attribute has been filtered")
}

// AllAttributes returns every Attribute which hasn't been filtered.
func (v *InstancesView) AllAttributes() []Attribute {

	if v.attrs == nil {
		return v.src.AllAttributes()
	}

	ret := make([]Attribute, len(v.attrs))

	for i, a := range v.attrs {
		ret[i] = a.GetAttribute()
	}

	return ret
}

// AddClassAttribute adds the given Attribute to the set of defined
// class Attributes, if it hasn't been filtered.
func (v *InstancesView) AddClassAttribute(a Attribute) error {
	// Check that this Attribute is defined
	matched := false
	for _, r := range v.AllAttributes() {
		if r.Equals(a) {
			matched = true
		}
	}
	if !matched {
		return fmt.Errorf("Attribute has been filtered")
	}

	v.classAttrs[a] = true
	return nil
}

// RemoveClassAttribute removes the given Attribute from the set of
// class Attributes.
func (v *InstancesView) RemoveClassAttribute(a Attribute) error {
	v.classAttrs[a] = false
	return nil
}

// AllClassAttributes returns all the Attributes currently defined
// as being class Attributes.
func (v *InstancesView) AllClassAttributes() []Attribute {
	ret := make([]Attribute, 0)
	for a := range v.classAttrs {
		if v.classAttrs[a] {
			ret = append(ret, a)
		}
	}
	return ret
}

// Get returns a sequence of bytes stored under a given Attribute
// on a given row.
//
// IMPORTANT: The AttributeSpec is unverified, meaning it's possible
// to return values from Attributes filtered by this InstancesView
// if the underlying AttributeSpec is known.
func (v *InstancesView) Get(as AttributeSpec, row int) []byte {
	// Change the row if necessary
	row = v.resolveRow(row)
	if row == -1 {
		panic("Out of range")
	}
	return v.src.Get(as, row)
}

// MapOverRows, see DenseInstances.MapOverRows.
//
// IMPORTANT: MapOverRows is not guaranteed to be ordered, but this one
// especially so.
func (v *InstancesView) MapOverRows(as []AttributeSpec, rowFunc func([][]byte, int) (bool, error)) error {
	if v.maskRows {
		rowBuf := make([][]byte, len(as))
		for r := range v.rows {
			row := v.rows[r]
			for i, a := range as {
				rowBuf[i] = v.src.Get(a, row)
			}
			ok, err := rowFunc(rowBuf, r)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
		}
		return nil
	} else {
		return v.src.MapOverRows(as, rowFunc)
	}
}

// Size Returns the number of Attributes and rows this InstancesView
// contains.
func (v *InstancesView) Size() (int, int) {
	// Get the original size
	hSize, vSize := v.src.Size()
	// Adjust to the number of defined Attributes
	if v.attrs != nil {
		hSize = len(v.attrs)
	}
	// Adjust to the number of defined rows
	if v.rows != nil {
		if v.maskRows {
			vSize = len(v.rows)
		} else if len(v.rows) > vSize {
			vSize = len(v.rows)
		}
	}
	return hSize, vSize
}

// String returns a human-readable summary of this InstancesView.
func (v *InstancesView) String() string {
	var buffer bytes.Buffer
	maxRows := 30

	// Get all Attribute information
	as := ResolveAllAttributes(v)

	// Print header
	cols, rows := v.Size()
	buffer.WriteString("InstancesView with ")
	buffer.WriteString(fmt.Sprintf("%d row(s) ", rows))
	buffer.WriteString(fmt.Sprintf("%d attribute(s)\n", cols))
	if v.attrs != nil {
		buffer.WriteString(fmt.Sprintf("With defined Attribute view\n"))
	}
	if v.rows != nil {
		buffer.WriteString(fmt.Sprintf("With defined Row view\n"))
	}
	if v.maskRows {
		buffer.WriteString("Row masking on.\n")
	}
	buffer.WriteString(fmt.Sprintf("Attributes:\n"))

	for _, a := range as {
		prefix := "\t"
		if v.classAttrs[a.attr] {
			prefix = "*\t"
		}
		buffer.WriteString(fmt.Sprintf("%s%s\n", prefix, a.attr))
	}

	// Print data
	if rows < maxRows {
		maxRows = rows
	}
	buffer.WriteString("Data:")
	for i := 0; i < maxRows; i++ {
		buffer.WriteString("\t")
		for _, a := range as {
			val := v.Get(a, i)
			buffer.WriteString(fmt.Sprintf("%s ", a.attr.GetStringFromSysVal(val)))
		}
		buffer.WriteString("\n")
	}

	missingRows := rows - maxRows
	if missingRows != 0 {
		buffer.WriteString(fmt.Sprintf("\t...\n%d row(s) undisplayed", missingRows))
	} else {
		buffer.WriteString("All rows displayed")
	}

	return buffer.String()
}

// RowString returns a string representation of a given row.
func (v *InstancesView) RowString(row int) string {
	var buffer bytes.Buffer
	as := ResolveAllAttributes(v)
	first := true
	for _, a := range as {
		val := v.Get(a, row)
		prefix := " "
		if first {
			prefix = ""
			first = false
		}
		buffer.WriteString(fmt.Sprintf("%s%s", prefix, a.attr.GetStringFromSysVal(val)))
	}
	return buffer.String()
}
