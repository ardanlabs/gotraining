package base

import (
	"os"
)

// ParseCSVGetRows returns the number of rows in a given file.
func ParseCSVGetRows(filepath string) (int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return ParseCSVGetRowsFromReader(f)
}

// ParseCSVEstimateFilePrecision determines what the maximum number of
// digits occuring anywhere after the decimal point within the file.
func ParseCSVEstimateFilePrecision(filepath string) (int, error) {
	// Open the source file
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return ParseCSVEstimateFilePrecisionFromReader(f)
}

// ParseCSVGetAttributes returns an ordered slice of appropriate-ly typed
// and named Attributes.
func ParseCSVGetAttributes(filepath string, hasHeaders bool) []Attribute {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return ParseCSVGetAttributesFromReader(f, hasHeaders)
}

// ParseCSVSniffAttributeNames returns a slice containing the top row
// of a given CSV file, or placeholders if hasHeaders is false.
func ParseCSVSniffAttributeNames(filepath string, hasHeaders bool) []string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return ParseCSVSniffAttributeNamesFromReader(f, hasHeaders)
}

// ParseCSVSniffAttributeTypes returns a slice of appropriately-typed Attributes.
//
// The type of a given attribute is determined by looking at the first data row
// of the CSV.
func ParseCSVSniffAttributeTypes(filepath string, hasHeaders bool) []Attribute {
	// Open file
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return ParseCSVSniffAttributeTypesFromReader(f, hasHeaders)
}

// ParseCSVToInstances reads the CSV file given by filepath and returns
// the read Instances.
func ParseCSVToInstances(filepath string, hasHeaders bool) (instances *DenseInstances, err error) {
	// Open the file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseCSVToInstancesFromReader(f, hasHeaders)
}

// ParseCSVToInstancesTemplated reads the CSV file given by filepath and returns
// the read Instances, using another already read DenseInstances as a template.
func ParseCSVToTemplatedInstances(filepath string, hasHeaders bool, template *DenseInstances) (instances *DenseInstances, err error) {
	// Open the file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseCSVToTemplatedInstancesFromReader(f, hasHeaders, template)
}

// ParseCSVToInstancesWithAttributeGroups reads the CSV file given by filepath,
// and returns the read DenseInstances, but also makes sure to group any Attributes
// specified in the first argument and also any class Attributes specified in the second
func ParseCSVToInstancesWithAttributeGroups(filepath string, attrGroups, classAttrGroups map[string]string, attrOverrides map[int]Attribute, hasHeaders bool) (instances *DenseInstances, err error) {
	// Open file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseCSVToInstancesWithAttributeGroupsFromReader(f, attrGroups, classAttrGroups, attrOverrides, hasHeaders)
}
