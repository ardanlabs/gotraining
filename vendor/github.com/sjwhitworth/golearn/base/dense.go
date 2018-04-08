package base

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

// DenseInstances stores each Attribute value explicitly
// in a large grid.
type DenseInstances struct {
	agMap        map[string]int
	agRevMap     map[int]string
	ags          []AttributeGroup
	lock         sync.Mutex
	fixed        bool
	classAttrs   map[AttributeSpec]bool
	maxRow       int
	attributes   []Attribute
	tmpAttrAgMap map[Attribute]string
	// Counters for each AttributeGroup type
	floatRowSizeBytes int
	catRowSizeBytes   int
	binRowSizeBits    int
}

// NewDenseInstances generates a new DenseInstances set
// with an anonymous EDF mapping and default settings.
func NewDenseInstances() *DenseInstances {
	return &DenseInstances{
		make(map[string]int),
		make(map[int]string),
		make([]AttributeGroup, 0),
		sync.Mutex{},
		false,
		make(map[AttributeSpec]bool),
		0,
		make([]Attribute, 0),
		make(map[Attribute]string),
		0,
		0,
		0,
	}
}

func copyFixedDataGridStructure(of FixedDataGrid) (*DenseInstances, []AttributeSpec, []AttributeSpec) {
	ret := NewDenseInstances() // Create the skeleton
	// Attribute creation
	attrs := of.AllAttributes()
	specs1 := make([]AttributeSpec, len(attrs))
	specs2 := make([]AttributeSpec, len(attrs))
	for i, a := range attrs {
		// Retrieve old AttributeSpec
		s, err := of.GetAttribute(a)
		if err != nil {
			panic(err)
		}
		specs1[i] = s
		// Add and store new AttributeSpec
		specs2[i] = ret.AddAttribute(a)
	}

	// Add class attributes
	cAttrs := of.AllClassAttributes()
	for _, a := range cAttrs {
		ret.AddClassAttribute(a)
	}
	return ret, specs1, specs2
}

// NewStructuralCopy generates an empty DenseInstances with the same layout as
// an existing FixedDataGrid, but with no data.
func NewStructuralCopy(of FixedDataGrid) *DenseInstances {
	ret, _, _ := copyFixedDataGridStructure(of)
	return ret
}

// NewDenseCopy generates a new DenseInstances set
// from an existing FixedDataGrid.
func NewDenseCopy(of FixedDataGrid) *DenseInstances {

	ret, specs1, specs2 := copyFixedDataGridStructure(of)

	// Allocate memory
	_, rows := of.Size()
	ret.Extend(rows)

	// Copy each row from the old one to the new
	of.MapOverRows(specs1, func(v [][]byte, r int) (bool, error) {
		for i, c := range v {
			ret.Set(specs2[i], r, c)
		}
		return true, nil
	})

	return ret
}

//
// AttributeGroup functions
//

// createAttributeGroup adds a new AttributeGroup to this set of Instances
// IMPORTANT: do not call unless you've acquired the lock
func (inst *DenseInstances) createAttributeGroup(name string, size int) {

	var agAdd AttributeGroup

	if inst.fixed {
		panic("Can't add additional Attributes")
	}

	// Create the AttributeGroup information
	if size != 0 {
		ag := new(FixedAttributeGroup)
		ag.parent = inst
		ag.attributes = make([]Attribute, 0)
		ag.size = size
		ag.alloc = make([]byte, 0)
		agAdd = ag
	} else {
		ag := new(BinaryAttributeGroup)
		ag.parent = inst
		ag.attributes = make([]Attribute, 0)
		ag.size = size
		ag.alloc = make([]byte, 0)
		agAdd = ag
	}
	inst.agMap[name] = len(inst.ags)
	inst.agRevMap[len(inst.ags)] = name
	inst.ags = append(inst.ags, agAdd)
}

// CreateAttributeGroup adds a new AttributeGroup to this set of instances
// with a given name. If the size is 0, a bit-ag is added
// if the size of not 0, then the size of each ag attribute
// is set to that number of bytes.
func (inst *DenseInstances) CreateAttributeGroup(name string, size int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("CreateAttributeGroup: %v (not created)", r)
			}
		}
	}()

	inst.lock.Lock()
	defer inst.lock.Unlock()

	inst.createAttributeGroup(name, size)
	return nil
}

// AllAttributeGroups returns a copy of the available AttributeGroups
func (inst *DenseInstances) AllAttributeGroups() map[string]AttributeGroup {
	ret := make(map[string]AttributeGroup)
	for a := range inst.agMap {
		ret[a] = inst.ags[inst.agMap[a]]
	}
	return ret
}

// GetAttributeGroup returns a reference to a AttributeGroup of a given name /
func (inst *DenseInstances) GetAttributeGroup(name string) (AttributeGroup, error) {
	inst.lock.Lock()
	defer inst.lock.Unlock()

	// Check if the ag exists
	if id, ok := inst.agMap[name]; !ok {
		return nil, fmt.Errorf("AttributeGroup '%s' doesn't exist", name)
	} else {
		// Return the ag
		return inst.ags[id], nil
	}
}

//
// Attribute creation and handling
//

// AddAttribute adds an Attribute to this set of DenseInstances
// Creates a default AttributeGroup for it if a suitable one doesn't exist.
// Returns an AttributeSpec for subsequent Set() calls.
//
// IMPORTANT: will panic if storage has been allocated via Extend.
func (inst *DenseInstances) AddAttribute(a Attribute) AttributeSpec {
	var ok bool
	inst.lock.Lock()
	defer inst.lock.Unlock()

	if inst.fixed {
		panic("Can't add additional Attributes")
	}

	cur := 0
	// Generate a default AttributeGroup name
	ag := "FLOAT"
	generatingBinClass := false
	if ag, ok = inst.tmpAttrAgMap[a]; ok {
		// Retrieved the group id
	} else if _, ok := a.(*CategoricalAttribute); ok {
		inst.catRowSizeBytes += 8
		cur = inst.catRowSizeBytes / os.Getpagesize()
		ag = fmt.Sprintf("CAT%d", cur)
	} else if _, ok := a.(*FloatAttribute); ok {
		inst.floatRowSizeBytes += 8
		cur = inst.floatRowSizeBytes / os.Getpagesize()
		ag = fmt.Sprintf("FLOAT%d", cur)
	} else if _, ok := a.(*BinaryAttribute); ok {
		inst.binRowSizeBits++
		cur = (inst.binRowSizeBits / 8) / os.Getpagesize()
		ag = fmt.Sprintf("BIN%d", cur)
		generatingBinClass = true
	} else {
		panic("Unrecognised Attribute type")
	}

	// Create the ag if it doesn't exist
	if _, ok := inst.agMap[ag]; !ok {
		if !generatingBinClass {
			inst.createAttributeGroup(ag, 8)
		} else {
			inst.createAttributeGroup(ag, 0)
		}
	}
	id := inst.agMap[ag]
	p := inst.ags[id]
	p.AddAttribute(a)
	inst.attributes = append(inst.attributes, a)
	return AttributeSpec{id, len(p.Attributes()) - 1, a}
}

// AddAttributeToAttributeGroup adds an Attribute to a given ag
func (inst *DenseInstances) AddAttributeToAttributeGroup(newAttribute Attribute, ag string) (AttributeSpec, error) {
	inst.lock.Lock()
	defer inst.lock.Unlock()

	// Check if the ag exists
	if _, ok := inst.agMap[ag]; !ok {
		return AttributeSpec{-1, 0, nil}, fmt.Errorf("AttributeGroup '%s' doesn't exist. Call CreateAttributeGroup() first", ag)
	}

	id := inst.agMap[ag]
	p := inst.ags[id]
	for i, a := range p.Attributes() {
		if !a.Compatible(newAttribute) {
			return AttributeSpec{-1, 0, nil}, fmt.Errorf("Attribute %s is not Compatible with %s in pond '%s' (position %d)", newAttribute, a, ag, i)
		}
	}

	p.AddAttribute(newAttribute)
	inst.attributes = append(inst.attributes, newAttribute)
	return AttributeSpec{id, len(p.Attributes()) - 1, newAttribute}, nil
}

// GetAttribute returns an Attribute equal to the argument.
//
// TODO: Write a function to pre-compute this once we've allocated
// TODO: Write a utility function which retrieves all AttributeSpecs for
// a given instance set.
func (inst *DenseInstances) GetAttribute(get Attribute) (AttributeSpec, error) {
	inst.lock.Lock()
	defer inst.lock.Unlock()

	for i, p := range inst.ags {
		for j, a := range p.Attributes() {
			if a.Equals(get) {
				return AttributeSpec{i, j, a}, nil
			}
		}
	}

	return AttributeSpec{-1, 0, nil}, fmt.Errorf("Couldn't resolve %s", get)
}

// AllAttributes returns a slice of all Attributes.
func (inst *DenseInstances) AllAttributes() []Attribute {
	inst.lock.Lock()
	defer inst.lock.Unlock()

	ret := make([]Attribute, 0)
	for _, p := range inst.ags {
		for _, a := range p.Attributes() {
			ret = append(ret, a)
		}
	}

	return ret
}

// AddClassAttribute sets an Attribute to be a class Attribute.
func (inst *DenseInstances) AddClassAttribute(a Attribute) error {

	as, err := inst.GetAttribute(a)
	if err != nil {
		return err
	}

	inst.lock.Lock()
	defer inst.lock.Unlock()

	inst.classAttrs[as] = true
	return nil
}

// RemoveClassAttribute removes an Attribute from the set of class Attributes.
func (inst *DenseInstances) RemoveClassAttribute(a Attribute) error {

	as, err := inst.GetAttribute(a)
	if err != nil {
		return err
	}

	inst.lock.Lock()
	defer inst.lock.Unlock()

	inst.classAttrs[as] = false
	return nil
}

// AllClassAttributes returns a slice of Attributes which have
// been designated class Attributes.
func (inst *DenseInstances) AllClassAttributes() []Attribute {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	return inst.allClassAttributes()
}

// allClassAttributes returns a slice of Attributes which have
// been designated class Attributes (doesn't lock)
func (inst *DenseInstances) allClassAttributes() []Attribute {
	var ret []Attribute
	for a := range inst.classAttrs {
		if inst.classAttrs[a] {
			ret = append(ret, a.attr)
		}
	}
	return ret
}

//
// Allocation functions
//

// realiseAttributeGroups decides which Attributes are going
// to be stored in which groups
func (inst *DenseInstances) realiseAttributeGroups() error {
	for a := range inst.tmpAttrAgMap {
		// Generate a default AttributeGroup name
		ag := inst.tmpAttrAgMap[a]

		// Augment with some additional information
		// Find out whether this attribute is also a class
		classAttribute := false
		for _, c := range inst.allClassAttributes() {
			if c.Equals(a) {
				classAttribute = true
			}
		}

		// This might result in multiple ClassAttribute groups
		// but hopefully nothing too crazy
		if classAttribute {
			// ag = fmt.Sprintf("CLASS_%s", ag)
		}

		// Create the ag if it doesn't exist
		if agId, ok := inst.agMap[ag]; !ok {
			_, generatingBinClass := inst.ags[agId].(*BinaryAttributeGroup)
			if !generatingBinClass {
				inst.createAttributeGroup(ag, 8)
			} else {
				inst.createAttributeGroup(ag, 0)
			}
		}
		id := inst.agMap[ag]
		p := inst.ags[id]
		err := p.AddAttribute(a)
		if err != nil {
			return err
		}
	}
	return nil
}

// Extend extends this set of Instances to store rows additional rows.
// It's recommended to set rows to something quite large.
//
// IMPORTANT: panics if the allocation fails
func (inst *DenseInstances) Extend(rows int) error {

	inst.lock.Lock()
	defer inst.lock.Unlock()

	if !inst.fixed {
		err := inst.realiseAttributeGroups()
		if err != nil {
			return err
		}
	}

	for _, p := range inst.ags {

		// Compute ag row storage requirements
		rowSize := p.RowSizeInBytes()

		// How many bytes?
		allocSize := rows * rowSize

		p.resize(allocSize)

	}
	inst.fixed = true
	inst.maxRow += rows
	return nil
}

// Set sets a particular Attribute (given as an AttributeSpec) on a particular
// row to a particular value.
//
// AttributeSpecs can be obtained using GetAttribute() or AddAttribute().
//
// IMPORTANT: Will panic() if the AttributeSpec isn't valid
//
// IMPORTANT: Will panic() if the row is too large
//
// IMPORTANT: Will panic() if the val is not the right length
func (inst *DenseInstances) Set(a AttributeSpec, row int, val []byte) {
	inst.ags[a.pond].set(a.position, row, val)
}

// Get gets a particular Attribute (given as an AttributeSpec) on a particular
// row.
// AttributeSpecs can be obtained using GetAttribute() or AddAttribute()
func (inst *DenseInstances) Get(a AttributeSpec, row int) []byte {
	return inst.ags[a.pond].get(a.position, row)
}

// RowString returns a string representation of a given row.
func (inst *DenseInstances) RowString(row int) string {
	var buffer bytes.Buffer
	first := true
	for _, p := range inst.ags {
		if first {
			first = false
		} else {
			buffer.WriteString(" ")
		}
		p.appendToRowBuf(row, &buffer)
	}
	return buffer.String()
}

// MapOverRows passes each row map into a function.
// First argument is a list of AttributeSpec in the order
// they're needed in for the function. The second is the function
// to call on each row.
func (inst *DenseInstances) MapOverRows(asv []AttributeSpec, mapFunc func([][]byte, int) (bool, error)) error {
	rowBuf := make([][]byte, len(asv))
	for i := 0; i < inst.maxRow; i++ {
		for j, as := range asv {
			p := inst.ags[as.pond]
			rowBuf[j] = p.get(as.position, i)
		}
		ok, err := mapFunc(rowBuf, i)
		if err != nil {
			return err
		}
		if !ok {
			break
		}
	}
	return nil
}

// Size returns the number of Attributes as the first return value
// and the maximum allocated row as the second value.
func (inst *DenseInstances) Size() (int, int) {
	return len(inst.AllAttributes()), inst.maxRow
}

// swapRows swaps over rows i and j
func (inst *DenseInstances) swapRows(i, j int) {
	as := ResolveAllAttributes(inst)
	for _, a := range as {
		v1 := inst.Get(a, i)
		v2 := inst.Get(a, j)
		v3 := make([]byte, len(v2))
		copy(v3, v2)
		inst.Set(a, j, v1)
		inst.Set(a, i, v3)
	}
}

// String returns a human-readable summary of this dataset.
func (inst *DenseInstances) String() string {
	var buffer bytes.Buffer

	// Get all Attribute information
	as := ResolveAllAttributes(inst)

	// Print header
	cols, rows := inst.Size()
	buffer.WriteString("Instances with ")
	buffer.WriteString(fmt.Sprintf("%d row(s) ", rows))
	buffer.WriteString(fmt.Sprintf("%d attribute(s)\n", cols))
	buffer.WriteString(fmt.Sprintf("Attributes: \n"))

	for _, a := range as {
		prefix := "\t"
		if inst.classAttrs[a] {
			prefix = "*\t"
		}
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
			val := inst.Get(a, i)
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
