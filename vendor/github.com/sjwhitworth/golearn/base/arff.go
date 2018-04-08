package base

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

// SerializeInstancesToDenseARFF writes the given FixedDataGrid to a
// densely-formatted ARFF file.
func SerializeInstancesToDenseARFF(inst FixedDataGrid, path, relation string) error {

	// Get all of the Attributes in a reasonable order
	attrs := NonClassAttributes(inst)
	cAttrs := inst.AllClassAttributes()
	for _, c := range cAttrs {
		attrs = append(attrs, c)
	}

	return SerializeInstancesToDenseARFFWithAttributes(inst, attrs, path, relation)

}

// SerializeInstancesToDenseARFFWithAttributes writes the given FixedDataGrid to a
// densely-formatted ARFF file with the header Attributes in the order given.
func SerializeInstancesToDenseARFFWithAttributes(inst FixedDataGrid, rawAttrs []Attribute, path, relation string) error {

	// Open output file
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return SerializeInstancesToWriterDenseARFFWithAttributes(f, inst, rawAttrs, relation)

}

func SerializeInstancesToWriterDenseARFFWithAttributes(w io.Writer, inst FixedDataGrid, rawAttrs []Attribute, relation string) error {
	// Write @relation header
	fmt.Fprintf(w, "@relation %s\n\n", relation)

	// Get all Attribute specifications
	attrs := ResolveAttributes(inst, rawAttrs)

	// Write Attribute information
	for _, s := range attrs {
		attr := s.attr
		t := "real"
		if a, ok := attr.(*CategoricalAttribute); ok {
			vals := a.GetValues()
			t = fmt.Sprintf("{%s}", strings.Join(vals, ", "))
		}
		fmt.Fprintf(w, "@attribute %s %s\n", attr.GetName(), t)
	}
	fmt.Fprint(w, "\n@data\n")

	buf := make([]string, len(attrs))
	inst.MapOverRows(attrs, func(val [][]byte, row int) (bool, error) {
		for i, v := range val {
			buf[i] = attrs[i].attr.GetStringFromSysVal(v)
		}
		fmt.Fprint(w, strings.Join(buf, ","))
		fmt.Fprint(w, "\n")
		return true, nil
	})

	return nil
}

// ParseARFFGetRows returns the number of data rows in an ARFF file.
func ParseARFFGetRows(filepath string) (int, error) {

	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	counting := false
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if counting {
			if line[0] == '@' {
				continue
			}
			if line[0] == '%' {
				continue
			}
			count++
			continue
		}
		if line[0] == '@' {
			line = strings.ToLower(line)
			if line == "@data" {
				counting = true
			}
		}
	}
	return count, nil
}

// ParseARFFGetAttributes returns the set of Attributes represented in this ARFF
func ParseARFFGetAttributes(filepath string) []Attribute {
	var ret []Attribute

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var attr Attribute
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] != '@' {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		fields[0] = strings.ToLower(fields[0])
		attrType := strings.ToLower(fields[2])
		if fields[0] != "@attribute" {
			continue
		}
		switch attrType {
		case "real":
			attr = new(FloatAttribute)
			break
		default:
			if fields[2][0] == '{' {
				if strings.HasSuffix(fields[len(fields)-1], "}") {
					var cats []string
					if len(fields) > 3 {
						cats = fields[2:len(fields)]
					} else {
						cats = strings.Split(fields[2], ",")
					}
					if len(cats) == 0 {
						panic(fmt.Errorf("Empty categorical field on line '%s'", line))
					}
					cats[0] = cats[0][1:]                                            // Remove leading '{'
					cats[len(cats)-1] = cats[len(cats)-1][:len(cats[len(cats)-1])-1] // Remove trailing '}'
					for i, v := range cats {                                         // Miaow
						cats[i] = strings.TrimSpace(v)
						if strings.HasSuffix(cats[i], ",") {
							// Strip end comma
							cats[i] = cats[i][0 : len(cats[i])-1]
						}
					}
					attr = NewCategoricalAttribute()
					for _, v := range cats {
						attr.GetSysValFromString(v)
					}
				} else {
					panic(fmt.Errorf("Missing categorical bracket on line '%s'", line))
				}
			} else {
				panic(fmt.Errorf("Unsupported Attribute type %s on line '%s'", fields[2], line))
			}
		}

		if attr == nil {
			panic(fmt.Errorf(line))
		}
		attr.SetName(fields[1])
		ret = append(ret, attr)
	}

	maxPrecision, err := ParseCSVEstimateFilePrecision(filepath)
	if err != nil {
		panic(err)
	}
	for _, a := range ret {
		if f, ok := a.(*FloatAttribute); ok {
			f.Precision = maxPrecision
		}
	}
	return ret
}

// ParseDenseARFFBuildInstancesFromReader updates an [[#UpdatableDataGrid]] from a io.Reader
func ParseDenseARFFBuildInstancesFromReader(r io.Reader, attrs []Attribute, u UpdatableDataGrid) (err error) {
	var rowCounter int

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(err)
			}
			err = fmt.Errorf("Error at line %d (error %s)", rowCounter, r.(error))
		}
	}()

	scanner := bufio.NewScanner(r)
	reading := false
	specs := ResolveAttributes(u, attrs)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "%") {
			continue
		}
		if reading {
			buf := bytes.NewBuffer([]byte(line))
			reader := csv.NewReader(buf)
			for {
				r, err := reader.Read()
				if err == io.EOF {
					break
				} else if err != nil {
					return err
				}
				for i, v := range r {
					v = strings.TrimSpace(v)
					if a, ok := specs[i].attr.(*CategoricalAttribute); ok {
						if val := a.GetSysVal(v); val == nil {
							panic(fmt.Errorf("Unexpected class on line '%s'", line))
						}
					}
					u.Set(specs[i], rowCounter, specs[i].attr.GetSysValFromString(v))
				}
				rowCounter++
			}
		} else {
			line = strings.ToLower(line)
			line = strings.TrimSpace(line)
			if line == "@data" {
				reading = true
			}
		}
	}

	return nil
}

// ParseDenseARFFToInstances parses the dense ARFF File into a FixedDataGrid
func ParseDenseARFFToInstances(filepath string) (ret *DenseInstances, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	// Find the number of rows in the file
	rows, err := ParseARFFGetRows(filepath)
	if err != nil {
		return nil, err
	}

	// Get the Attributes we want
	attrs := ParseARFFGetAttributes(filepath)

	// Allocate return value
	ret = NewDenseInstances()

	// Add all the Attributes
	for _, a := range attrs {
		ret.AddAttribute(a)
	}

	// Set the last Attribute as the class
	ret.AddClassAttribute(attrs[len(attrs)-1])
	ret.Extend(rows)

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read the data
	// Seek past the header
	err = ParseDenseARFFBuildInstancesFromReader(f, attrs, ret)
	if err != nil {
		ret = nil
	}
	return ret, err
}
