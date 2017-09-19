package base

import (
	"bytes"
	"fmt"
)

// Maybe included a TransformedAttribute struct
// so we can map from ClassAttribute to ClassAttribute

// LazilyFilteredInstances map a Filter over an underlying
// FixedDataGrid and are a memory-efficient way of applying them.
type LazilyFilteredInstances struct {
	filter        Filter
	src           FixedDataGrid
	attrs         []FilteredAttribute
	classAttrs    map[Attribute]bool
	unfilteredMap map[Attribute]bool
}

// NewLazilyFitleredInstances returns a new FixedDataGrid after
// applying the given Filter to the Attributes it includes. Unfiltered
// Attributes are passed through without modification.
func NewLazilyFilteredInstances(src FixedDataGrid, f Filter) *LazilyFilteredInstances {

	// Get the Attributes after filtering
	attrs := f.GetAttributesAfterFiltering()

	// Build a set of Attributes which have undergone filtering
	unFilteredMap := make(map[Attribute]bool)
	for _, a := range src.AllAttributes() {
		unFilteredMap[a] = true
	}
	for _, a := range attrs {
		unFilteredMap[a.Old] = false
	}

	// Create the return structure
	ret := &LazilyFilteredInstances{
		f,
		src,
		attrs,
		make(map[Attribute]bool),
		unFilteredMap,
	}

	// Transfer class Attributes
	for _, a := range src.AllClassAttributes() {
		ret.AddClassAttribute(a)
	}
	return ret
}

// GetAttribute returns an AttributeSpecification for a given Attribute
func (l *LazilyFilteredInstances) GetAttribute(target Attribute) (AttributeSpec, error) {
	if l.unfilteredMap[target] {
		return l.src.GetAttribute(target)
	}
	var ret AttributeSpec
	ret.pond = -1
	for i, a := range l.attrs {
		if a.New.Equals(target) {
			ret.position = i
			ret.attr = target
			return ret, nil
		}
	}
	return ret, fmt.Errorf("Couldn't resolve %s", target)
}

// AllAttributes returns every Attribute defined in the source datagrid,
// in addition to the revised Attributes created by the filter.
func (l *LazilyFilteredInstances) AllAttributes() []Attribute {
	ret := make([]Attribute, 0)
	for _, a := range l.src.AllAttributes() {
		if l.unfilteredMap[a] {
			ret = append(ret, a)
		} else {
			for _, b := range l.attrs {
				if a.Equals(b.Old) {
					ret = append(ret, b.New)
				}
			}
		}
	}
	return ret
}

// AddClassAttribute adds a given Attribute (either before or after filtering)
// to the set of defined class Attributes.
func (l *LazilyFilteredInstances) AddClassAttribute(cls Attribute) error {
	if l.unfilteredMap[cls] {
		l.classAttrs[cls] = true
		return nil
	}
	matched := false
	for _, a := range l.attrs {
		if a.Old.Equals(cls) || a.New.Equals(cls) {
			l.classAttrs[a.New] = true
			matched = true
		}
	}
	if !matched {
		return fmt.Errorf("Attribute %s could not be resolved", cls)
	}
	return nil
}

// RemoveClassAttribute removes a given Attribute (either before or
// after filtering) from the set of defined class Attributes.
func (l *LazilyFilteredInstances) RemoveClassAttribute(cls Attribute) error {
	if l.unfilteredMap[cls] {
		l.classAttrs[cls] = false
		return nil
	}
	for _, a := range l.attrs {
		if a.Old.Equals(cls) || a.New.Equals(cls) {
			l.classAttrs[a.New] = false
			return nil
		}
	}
	return fmt.Errorf("Attribute %s could not be resolved", cls)
}

// AllClassAttributes returns details of all Attributes currently specified
// as being class Attributes.
//
// If applicable, the Attributes returned are those after modification
// by the Filter.
func (l *LazilyFilteredInstances) AllClassAttributes() []Attribute {
	ret := make([]Attribute, 0)
	for a := range l.classAttrs {
		if l.classAttrs[a] {
			ret = append(ret, a)
		}
	}
	return ret
}

func (l *LazilyFilteredInstances) transformNewToOldAttribute(as AttributeSpec) (AttributeSpec, error) {
	if l.unfilteredMap[as.GetAttribute()] {
		return as, nil
	}
	for _, a := range l.attrs {
		if a.Old.Equals(as.attr) || a.New.Equals(as.attr) {
			as, err := l.src.GetAttribute(a.Old)
			if err != nil {
				return AttributeSpec{}, fmt.Errorf("Internal error in Attribute resolution: '%s'", err)
			}
			return as, nil
		}
	}
	return AttributeSpec{}, fmt.Errorf("No matching Attribute")
}

// Get returns a transformed byte slice stored at a given AttributeSpec and row.
func (l *LazilyFilteredInstances) Get(as AttributeSpec, row int) []byte {
	asOld, err := l.transformNewToOldAttribute(as)
	if err != nil {
		panic(fmt.Sprintf("Attribute %s could not be resolved. (Error: %s)", as.String(), err.Error()))
	}
	byteSeq := l.src.Get(asOld, row)
	if l.unfilteredMap[as.attr] {
		return byteSeq
	}
	newByteSeq := l.filter.Transform(asOld.attr, as.attr, byteSeq)
	return newByteSeq
}

// MapOverRows maps an iteration mapFunc over the bytes contained in the source
// FixedDataGrid, after modification by the filter.
func (l *LazilyFilteredInstances) MapOverRows(asv []AttributeSpec, mapFunc func([][]byte, int) (bool, error)) error {

	// Have to transform each item of asv into an
	// AttributeSpec in the original
	oldAsv := make([]AttributeSpec, len(asv))
	for i, a := range asv {
		old, err := l.transformNewToOldAttribute(a)
		if err != nil {
			return fmt.Errorf("Couldn't fetch old Attribute: '%s'", a.String())
		}
		oldAsv[i] = old
	}

	// Then map over each row in the original
	newRowBuf := make([][]byte, len(asv))
	return l.src.MapOverRows(oldAsv, func(oldRow [][]byte, oldRowNo int) (bool, error) {
		for i, b := range oldRow {
			newField := l.filter.Transform(oldAsv[i].attr, asv[i].attr, b)
			newRowBuf[i] = newField
		}
		return mapFunc(newRowBuf, oldRowNo)
	})
}

// RowString returns a string representation of a given row
// after filtering.
func (l *LazilyFilteredInstances) RowString(row int) string {
	var buffer bytes.Buffer

	as := ResolveAllAttributes(l) // Retrieve all Attribute data
	first := true                 // Decide whether to prefix

	for _, a := range as {
		prefix := " " // What to print before value
		if first {
			first = false // Don't print space on first value
			prefix = ""
		}
		val := l.Get(a, row) // Retrieve filtered value
		buffer.WriteString(fmt.Sprintf("%s%s", prefix, a.attr.GetStringFromSysVal(val)))
	}

	return buffer.String() // Return the result
}

// Size returns the number of Attributes and rows of the underlying
// FixedDataGrid.
func (l *LazilyFilteredInstances) Size() (int, int) {
	return l.src.Size()
}

// String returns a human-readable summary of this FixedDataGrid
// after filtering.
func (l *LazilyFilteredInstances) String() string {
	var buffer bytes.Buffer

	// Decide on rows to print
	_, rows := l.Size()
	maxRows := 5
	if rows < maxRows {
		maxRows = rows
	}

	// Get all Attribute information
	as := ResolveAllAttributes(l)

	// Print header
	buffer.WriteString("Lazily filtered instances using ")
	buffer.WriteString(fmt.Sprintf("%s\n", l.filter))
	buffer.WriteString(fmt.Sprintf("Attributes: \n"))

	for _, a := range as {
		prefix := "\t"
		if l.classAttrs[a.attr] {
			prefix = "*\t"
		}
		buffer.WriteString(fmt.Sprintf("%s%s\n", prefix, a.attr))
	}

	buffer.WriteString("\nData:\n")
	for i := 0; i < maxRows; i++ {
		buffer.WriteString("\t")
		for _, a := range as {
			val := l.Get(a, i)
			buffer.WriteString(fmt.Sprintf("%s ", a.attr.GetStringFromSysVal(val)))
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}
