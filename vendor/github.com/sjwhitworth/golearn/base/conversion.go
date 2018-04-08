package base

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func checkAllAttributesAreFloat(attrs []Attribute) error {
	// Check that all the attributes are float
	for _, a := range attrs {
		if _, ok := a.(*FloatAttribute); !ok {
			return fmt.Errorf("All []Attributes to this method must be FloatAttributes")
		}
	}
	return nil
}

// ConvertRowToMat64 takes a list of Attributes, a FixedDataGrid
// and a row number, and returns the float values of that row
// in a mat.Dense format.
func ConvertRowToMat64(attrs []Attribute, f FixedDataGrid, r int) (*mat.Dense, error) {

	err := checkAllAttributesAreFloat(attrs)
	if err != nil {
		return nil, err
	}

	// Allocate the return value
	ret := mat.NewDense(1, len(attrs), nil)

	// Resolve all the attributes
	attrSpecs := ResolveAttributes(f, attrs)

	// Get the results
	for i, a := range attrSpecs {
		ret.Set(0, i, UnpackBytesToFloat(f.Get(a, r)))
	}

	// Return the result
	return ret, nil
}

// ConvertAllRowsToMat64 takes a list of Attributes and returns a vector
// of all rows in a mat.Dense format.
func ConvertAllRowsToMat64(attrs []Attribute, f FixedDataGrid) ([]*mat.Dense, error) {

	// Check for floats
	err := checkAllAttributesAreFloat(attrs)
	if err != nil {
		return nil, err
	}

	// Return value
	_, rows := f.Size()
	ret := make([]*mat.Dense, rows)

	// Resolve all attributes
	attrSpecs := ResolveAttributes(f, attrs)

	// Set the values in each return value
	for i := 0; i < rows; i++ {
		cur := mat.NewDense(1, len(attrs), nil)
		for j, a := range attrSpecs {
			cur.Set(0, j, UnpackBytesToFloat(f.Get(a, i)))
		}
		ret[i] = cur
	}
	return ret, nil
}
