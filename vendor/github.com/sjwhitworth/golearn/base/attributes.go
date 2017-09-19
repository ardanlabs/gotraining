package base

import (
	"encoding/json"
)

const (
	// CategoricalType is for Attributes which represent values distinctly.
	CategoricalType = iota
	// Float64Type should be replaced with a FractionalNumeric type [DEPRECATED].
	Float64Type
	BinaryType
)

// Attributes disambiguate columns of the feature matrix and declare their types.
type Attribute interface {
	json.Unmarshaler
	json.Marshaler
	// Returns the general characterstics of this Attribute .
	// to avoid the overhead of casting
	GetType() int
	// Returns the human-readable name of this Attribute.
	GetName() string
	// Sets the human-readable name of this Attribute.
	SetName(string)
	// Gets a human-readable overview of this Attribute for debugging.
	String() string
	// Converts a value given from a human-readable string into a system
	// representation. For example, a CategoricalAttribute with values
	// ["iris-setosa", "iris-virginica"] would return the float64
	// representation of 0 when given "iris-setosa".
	GetSysValFromString(string) []byte
	// Converts a given value from a system representation into a human
	// representation. For example, a CategoricalAttribute with values
	// ["iris-setosa", "iris-viriginica"] might return "iris-setosa"
	// when given 0.0 as the argument.
	GetStringFromSysVal([]byte) string
	// Tests for equality with another Attribute. Other Attributes are
	// considered equal if:
	// * They have the same type (i.e. FloatAttribute <> CategoricalAttribute)
	// * They have the same name
	// * If applicable, they have the same categorical values (though not
	//   necessarily in the same order).
	Equals(Attribute) bool
	// Tests whether two Attributes can be represented in the same pond
	// i.e. they're the same size, and their byte order makes them meaningful
	// when considered together
	Compatible(Attribute) bool
}
