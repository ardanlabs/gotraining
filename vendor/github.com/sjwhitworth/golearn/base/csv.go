package base

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// ParseCSVGetRows returns the number of rows in a given file.
func ParseCSVGetRows(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	counter := 0
	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
		counter++
	}
	return counter, nil
}

// ParseCSVEstimateFilePrecision determines what the maximum number of
// digits occuring anywhere after the decimal point within the file.
func ParseCSVEstimateFilePrecision(filepath string) (int, error) {
	// Creat a basic regexp
	rexp := regexp.MustCompile("[0-9]+(.[0-9]+)?")

	// Open the source file
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Scan through the file line-by-line
	maxL := 0
	scanner := bufio.NewScanner(f)
	lineCount := 0
	for scanner.Scan() {
		if lineCount > 5 {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '@' {
			continue
		}
		if line[0] == '%' {
			continue
		}
		matches := rexp.FindAllString(line, -1)
		for _, m := range matches {
			p := strings.Split(m, ".")
			if len(p) == 2 {
				l := len(p[len(p)-1])
				if l > maxL {
					maxL = l
				}
			}
		}
		lineCount++
	}
	return maxL, nil
}

// ParseCSVGetAttributes returns an ordered slice of appropriate-ly typed
// and named Attributes.
func ParseCSVGetAttributes(filepath string, hasHeaders bool) []Attribute {
	attrs := ParseCSVSniffAttributeTypes(filepath, hasHeaders)
	names := ParseCSVSniffAttributeNames(filepath, hasHeaders)
	for i, attr := range attrs {
		attr.SetName(names[i])
	}
	return attrs
}

// ParseCSVSniffAttributeNames returns a slice containing the top row
// of a given CSV file, or placeholders if hasHeaders is false.
func ParseCSVSniffAttributeNames(filepath string, hasHeaders bool) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		panic(err)
	}

	if hasHeaders {
		for i, h := range headers {
			headers[i] = strings.TrimSpace(h)
		}
		return headers
	}

	for i := range headers {
		headers[i] = fmt.Sprintf("%d", i)
	}
	return headers

}

// ParseCSVSniffAttributeTypes returns a slice of appropriately-typed Attributes.
//
// The type of a given attribute is determined by looking at the first data row
// of the CSV.
func ParseCSVSniffAttributeTypes(filepath string, hasHeaders bool) []Attribute {
	var attrs []Attribute
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Create the CSV reader
	reader := csv.NewReader(file)
	if hasHeaders {
		// Skip the headers
		_, err := reader.Read()
		if err != nil {
			panic(err)
		}
	}
	// Read the first line of the file
	columns, err := reader.Read()
	if err != nil {
		panic(err)
	}

	for _, entry := range columns {
		// Match the Attribute type with regular expressions
		entry = strings.Trim(entry, " ")
		matched, err := regexp.MatchString("^[-+]?[0-9]*\\.?[0-9]+([eE][-+]?[0-9]+)?$", entry)
		if err != nil {
			panic(err)
		}
		if matched {
			attrs = append(attrs, NewFloatAttribute(""))
		} else {
			attrs = append(attrs, new(CategoricalAttribute))
		}
	}

	// Estimate file precision
	maxP, err := ParseCSVEstimateFilePrecision(filepath)
	if err != nil {
		panic(err)
	}
	for _, a := range attrs {
		if f, ok := a.(*FloatAttribute); ok {
			f.Precision = maxP
		}
	}

	return attrs
}

// ParseCSVBuildInstancesFromReader updates an [[#UpdatableDataGrid]] from a io.Reader
func ParseCSVBuildInstancesFromReader(r io.Reader, attrs []Attribute, hasHeader bool, u UpdatableDataGrid) (err error) {
	var rowCounter int

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(err)
			}
			err = fmt.Errorf("Error at line %d (error %s)", rowCounter, r.(error))
		}
	}()

	specs := ResolveAttributes(u, attrs)
	reader := csv.NewReader(r)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if rowCounter == 0 {
			if hasHeader {
				hasHeader = false
				continue
			}
		}
		for i, v := range record {
			u.Set(specs[i], rowCounter, specs[i].attr.GetSysValFromString(strings.TrimSpace(v)))
		}
		rowCounter++
	}

	return nil
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

	// Read the number of rows in the file
	rowCount, err := ParseCSVGetRows(filepath)
	if err != nil {
		return nil, err
	}

	if hasHeaders {
		rowCount--
	}

	// Read the row headers
	attrs := ParseCSVGetAttributes(filepath, hasHeaders)
	specs := make([]AttributeSpec, len(attrs))
	// Allocate the Instances to return
	instances = NewDenseInstances()
	for i, a := range attrs {
		spec := instances.AddAttribute(a)
		specs[i] = spec
	}
	instances.Extend(rowCount)

	err = ParseCSVBuildInstancesFromReader(f, attrs, hasHeaders, instances)
	if err != nil {
		return nil, err
	}

	instances.AddClassAttribute(attrs[len(attrs)-1])

	return instances, nil
}

// ParseUtilsMatchAttrs tries to match the set of Attributes read from one file with
// those read from another, and writes the matching Attributes back to the original set.
func ParseMatchAttributes(attrs, templateAttrs []Attribute) {
	for i, a := range attrs {
		for _, b := range templateAttrs {
			if a.Equals(b) {
				attrs[i] = b
			} else if a.GetName() == b.GetName() {
				attrs[i] = b
			}
		}
	}
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

	// Read the number of rows in the file
	rowCount, err := ParseCSVGetRows(filepath)
	if err != nil {
		return nil, err
	}

	if hasHeaders {
		rowCount--
	}

	// Read the row headers
	attrs := ParseCSVGetAttributes(filepath, hasHeaders)
	templateAttrs := template.AllAttributes()
	ParseMatchAttributes(attrs, templateAttrs)

	// Allocate the Instances to return
	instances = CopyDenseInstances(template, templateAttrs)
	instances.Extend(rowCount)

	err = ParseCSVBuildInstancesFromReader(f, attrs, hasHeaders, instances)
	if err != nil {
		return nil, err
	}

	for _, a := range template.AllClassAttributes() {
		err = instances.AddClassAttribute(a)
		if err != nil {
			return nil, err
		}
	}

	return instances, nil
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

	// Read row count
	rowCount, err := ParseCSVGetRows(filepath)
	if err != nil {
		return nil, err
	}

	// Read the row headers
	attrs := ParseCSVGetAttributes(filepath, hasHeaders)
	for i := range attrs {
		if a, ok := attrOverrides[i]; ok {
			attrs[i] = a
		}
	}

	specs := make([]AttributeSpec, len(attrs))
	// Allocate the Instances to return
	instances = NewDenseInstances()

	//
	// Create all AttributeGroups
	agsToCreate := make(map[string]int)
	combinedAgs := make(map[string]string)
	for a := range attrGroups {
		agsToCreate[attrGroups[a]] = 0
		combinedAgs[a] = attrGroups[a]
	}
	for a := range classAttrGroups {
		agsToCreate[classAttrGroups[a]] = 8
		combinedAgs[a] = classAttrGroups[a]
	}

	// Decide the sizes
	for _, a := range attrs {
		if ag, ok := combinedAgs[a.GetName()]; ok {
			if _, ok := a.(*BinaryAttribute); ok {
				agsToCreate[ag] = 0
			} else {
				agsToCreate[ag] = 8
			}
		}
	}

	// Create them
	for i := range agsToCreate {
		size := agsToCreate[i]
		err = instances.CreateAttributeGroup(i, size)
		if err != nil {
			panic(err)
		}
	}

	// Add the Attributes to them
	for i, a := range attrs {
		var spec AttributeSpec
		if ag, ok := combinedAgs[a.GetName()]; ok {
			spec, err = instances.AddAttributeToAttributeGroup(a, ag)
			if err != nil {
				panic(err)
			}
			specs[i] = spec
		} else {
			spec = instances.AddAttribute(a)
		}
		specs[i] = spec
		if _, ok := classAttrGroups[a.GetName()]; ok {
			err = instances.AddClassAttribute(a)
			if err != nil {
				panic(err)
			}
		}
	}
	// Allocate
	instances.Extend(rowCount)

	err = ParseCSVBuildInstancesFromReader(f, attrs, hasHeaders, instances)
	if err != nil {
		return nil, err
	}

	return instances, nil

}
