package base

import (
	"bytes"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Mat64Instances struct {
	attributes []Attribute
	classAttrs map[int]bool
	Data       *mat.Dense
	rows       int
}

// InstancesFromMat64 returns a new Mat64Instances from a literal provided.
func InstancesFromMat64(rows, cols int, data *mat.Dense) *Mat64Instances {

	var ret Mat64Instances
	for i := 0; i < cols; i++ {
		ret.attributes = append(ret.attributes, NewFloatAttribute(fmt.Sprintf("%d", i)))
	}

	ret.classAttrs = make(map[int]bool)
	ret.Data = data
	ret.rows = rows

	return &ret
}

// GetAttribute returns an AttributeSpec from an Attribute field.
func (m *Mat64Instances) GetAttribute(a Attribute) (AttributeSpec, error) {
	for i, at := range m.attributes {
		if at.Equals(a) {
			return AttributeSpec{0, i, at}, nil
		}
	}
	return AttributeSpec{}, fmt.Errorf("Couldn't find a matching attribute")
}

// AllAttributes returns every defined Attribute.
func (m *Mat64Instances) AllAttributes() []Attribute {
	ret := make([]Attribute, len(m.attributes))
	for i, a := range m.attributes {
		ret[i] = a
	}
	return ret
}

// AddClassAttribute adds an attribute to the class set.
func (m *Mat64Instances) AddClassAttribute(a Attribute) error {
	as, err := m.GetAttribute(a)
	if err != nil {
		return err
	}

	m.classAttrs[as.position] = true
	return nil
}

// RemoveClassAttribute removes an attribute to the class set.
func (m *Mat64Instances) RemoveClassAttribute(a Attribute) error {
	as, err := m.GetAttribute(a)
	if err != nil {
		return err
	}

	m.classAttrs[as.position] = false
	return nil
}

// AllClassAttributes returns every class attribute.
func (m *Mat64Instances) AllClassAttributes() []Attribute {
	ret := make([]Attribute, 0)
	for i := range m.classAttrs {
		if m.classAttrs[i] {
			ret = append(ret, m.attributes[i])
		}
	}

	return ret
}

// Get returns the bytes at a given position
func (m *Mat64Instances) Get(as AttributeSpec, row int) []byte {
	val := m.Data.At(row, as.position)
	return PackFloatToBytes(val)
}

// MapOverRows is a convenience function for iteration
func (m *Mat64Instances) MapOverRows(as []AttributeSpec, f func([][]byte, int) (bool, error)) error {

	rowData := make([][]byte, len(as))
	for j, _ := range as {
		rowData[j] = make([]byte, 8)
	}
	for i := 0; i < m.rows; i++ {
		for j, as := range as {
			PackFloatToBytesInline(m.Data.At(i, as.position), rowData[j])
		}
		stat, err := f(rowData, i)
		if !stat {
			return err
		}
	}
	return nil
}

// RowString: should print the values of a row
// TODO: make this less half-assed
func (m *Mat64Instances) RowString(row int) string {
	return fmt.Sprintf("%d", row)
}

// Size returns the number of Attributes, then the number of rows
func (m *Mat64Instances) Size() (int, int) {
	return len(m.attributes), m.rows
}

// String returns a human-readable summary of this dataset.
func (m *Mat64Instances) String() string {
	var buffer bytes.Buffer

	// Get all Attribute information
	as := ResolveAllAttributes(m)

	// Print header
	cols, rows := m.Size()
	buffer.WriteString("Instances with ")
	buffer.WriteString(fmt.Sprintf("%d row(s) ", rows))
	buffer.WriteString(fmt.Sprintf("%d attribute(s)\n", cols))
	buffer.WriteString(fmt.Sprintf("Attributes: \n"))

	cnt := 0
	for _, a := range as {
		prefix := "\t"
		if m.classAttrs[cnt] {
			prefix = "*\t"
		}
		cnt++
		buffer.WriteString(fmt.Sprintf("%s%s\n", prefix, a.attr))
	}

	buffer.WriteString("\nData:\n")
	maxRows := 30
	if rows < maxRows {
		maxRows = rows
	}

	for i := 0; i < maxRows; i++ {
		buffer.WriteString("\t")
		for _, a := range as {
			val := m.Get(a, i)
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
