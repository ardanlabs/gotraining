package base

import (
	"encoding/json"
	"fmt"
)

// CategoricalAttribute is an Attribute implementation
// which stores discrete string values
// - useful for representing classes.
type CategoricalAttribute struct {
	Name   string
	values []string
}

// MarshalJSON returns a JSON version of this Attribute.
func (Attr *CategoricalAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": "categorical",
		"name": Attr.Name,
		"attr": map[string]interface{}{
			"values": Attr.values,
		},
	})
}

// UnmarshalJSON returns a JSON version of this Attribute.
func (Attr *CategoricalAttribute) UnmarshalJSON(data []byte) error {
	var d map[string]interface{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	for _, v := range d["values"].([]interface{}) {
		Attr.values = append(Attr.values, v.(string))
	}
	return nil
}

// NewCategoricalAttribute creates a blank CategoricalAttribute.
func NewCategoricalAttribute() *CategoricalAttribute {
	return &CategoricalAttribute{
		"",
		make([]string, 0),
	}
}

// GetValues returns all the values currently defined
func (Attr *CategoricalAttribute) GetValues() []string {
	return Attr.values
}

// GetName returns the human-readable name assigned to this attribute.
func (Attr *CategoricalAttribute) GetName() string {
	return Attr.Name
}

// SetName sets the human-readable name on this attribute.
func (Attr *CategoricalAttribute) SetName(name string) {
	Attr.Name = name
}

// GetType returns CategoricalType to avoid casting overhead.
func (Attr *CategoricalAttribute) GetType() int {
	return CategoricalType
}

// GetSysVal returns the system representation of userVal as an index into the Values slice
// If the userVal can't be found, it returns nothing.
func (Attr *CategoricalAttribute) GetSysVal(userVal string) []byte {
	for idx, val := range Attr.values {
		if val == userVal {
			return PackU64ToBytes(uint64(idx))
		}
	}
	return nil
}

// GetUsrVal returns a human-readable representation of the given sysVal.
//
// IMPORTANT: this function doesn't check the boundaries of the array.
func (Attr *CategoricalAttribute) GetUsrVal(sysVal []byte) string {
	idx := UnpackBytesToU64(sysVal)
	return Attr.values[idx]
}

// GetSysValFromString returns the system representation of rawVal
// as an index into the Values slice. If rawVal is not inside
// the Values slice, it is appended.
//
// IMPORTANT: If no system representation yet exists, this functions adds it.
// If you need to determine whether rawVal exists: use GetSysVal and check
// for a zero-length return value.
//
// Example: if the CategoricalAttribute contains the values ["iris-setosa",
// "iris-virginica"] and "iris-versicolor" is provided as the argument,
// the Values slide becomes ["iris-setosa", "iris-virginica", "iris-versicolor"]
// and 2.00 is returned as the system representation.
func (Attr *CategoricalAttribute) GetSysValFromString(rawVal string) []byte {
	// Match in raw values
	catIndex := -1
	for i, s := range Attr.values {
		if s == rawVal {
			catIndex = i
			break
		}
	}
	if catIndex == -1 {
		Attr.values = append(Attr.values, rawVal)
		catIndex = len(Attr.values) - 1
	}

	ret := PackU64ToBytes(uint64(catIndex))
	return ret
}

// String returns a human-readable summary of this Attribute.
//
// Returns a string containing the list of human-readable values this
// CategoricalAttribute can take.
func (Attr *CategoricalAttribute) String() string {
	return fmt.Sprintf("CategoricalAttribute(\"%s\", %s)", Attr.Name, Attr.values)
}

// GetStringFromSysVal returns a human-readable value from the given system-representation
// value val.
//
// IMPORTANT: This function calls panic() if the value is greater than
// the length of the array.
// TODO: Return a user-configurable default instead.
func (Attr *CategoricalAttribute) GetStringFromSysVal(rawVal []byte) string {
	convVal := int(UnpackBytesToU64(rawVal))
	if convVal >= len(Attr.values) {
		panic(fmt.Sprintf("Out of range: %d in %d (%s)", convVal, len(Attr.values), Attr))
	}
	return Attr.values[convVal]
}

// Equals checks equality against another Attribute.
//
// Two CategoricalAttributes are considered equal if they contain
// the same values and have the same name. Otherwise, this function
// returns false.
func (Attr *CategoricalAttribute) Equals(other Attribute) bool {
	attribute, ok := other.(*CategoricalAttribute)
	if !ok {
		// Not the same type, so can't be equal
		return false
	}
	if Attr.GetName() != attribute.GetName() {
		return false
	}

	// Check that this CategoricalAttribute has the same
	// values as the other, in the same order
	if len(attribute.values) != len(Attr.values) {
		return false
	}

	for i, a := range Attr.values {
		if a != attribute.values[i] {
			return false
		}
	}

	return true
}

// Compatible checks that this CategoricalAttribute has the same
// values as another, in the same order.
func (Attr *CategoricalAttribute) Compatible(other Attribute) bool {
	attribute, ok := other.(*CategoricalAttribute)
	if !ok {
		return false
	}

	// Check that this CategoricalAttribute has the same
	// values as the other, in the same order
	if len(attribute.values) != len(Attr.values) {
		return false
	}

	for i, a := range Attr.values {
		if a != attribute.values[i] {
			return false
		}
	}

	return true
}
