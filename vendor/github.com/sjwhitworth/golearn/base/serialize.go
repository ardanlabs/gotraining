package base

import (
	"archive/tar"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
)

const (
	SerializationFormatVersion = "golearn 0.5"
)

func SerializeInstancesToFile(inst FixedDataGrid, path string) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	err = SerializeInstances(inst, f)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return fmt.Errorf("Couldn't flush file: %s", err)
	}
	f.Close()
	return nil
}

func SerializeInstancesToCSV(inst FixedDataGrid, path string) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer func() {
		f.Sync()
		f.Close()
	}()

	return SerializeInstancesToCSVStream(inst, f)
}

func SerializeInstancesToCSVStream(inst FixedDataGrid, f io.Writer) error {
	// Create the CSV writer
	w := csv.NewWriter(f)

	colCount, _ := inst.Size()

	// Write out Attribute headers
	// Start with the regular Attributes
	normalAttrs := NonClassAttributes(inst)
	classAttrs := inst.AllClassAttributes()
	allAttrs := make([]Attribute, colCount)
	n := copy(allAttrs, normalAttrs)
	copy(allAttrs[n:], classAttrs)
	headerRow := make([]string, colCount)
	for i, v := range allAttrs {
		headerRow[i] = v.GetName()
	}
	w.Write(headerRow)

	specs := ResolveAttributes(inst, allAttrs)
	curRow := make([]string, colCount)
	inst.MapOverRows(specs, func(row [][]byte, rowNo int) (bool, error) {
		for i, v := range row {
			attr := allAttrs[i]
			curRow[i] = attr.GetStringFromSysVal(v)
		}
		w.Write(curRow)
		return true, nil
	})

	w.Flush()
	return nil
}

func writeAttributesToFilePart(attrs []Attribute, f *tar.Writer, name string) error {
	// Get the marshaled Attribute array
	body, err := json.Marshal(attrs)
	if err != nil {
		return err
	}

	// Write a header
	hdr := &tar.Header{
		Name: name,
		Size: int64(len(body)),
	}
	if err := f.WriteHeader(hdr); err != nil {
		return err
	}

	// Write the marshaled data
	if _, err := f.Write([]byte(body)); err != nil {
		return err
	}

	return nil
}

func getTarContent(tr *tar.Reader, name string) []byte {
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if hdr.Name == name {
			ret := make([]byte, hdr.Size)
			n, err := tr.Read(ret)
			if int64(n) != hdr.Size {
				panic("Size mismatch")
			}
			if err != nil {
				panic(err)
			}
			return ret
		}
	}
	panic("File not found!")
}

func deserializeAttributes(data []byte) []Attribute {

	// Define a JSON shim Attribute
	type JSONAttribute struct {
		Type string          `json:"type"`
		Name string          `json:"name"`
		Attr json.RawMessage `json:"attr"`
	}

	var ret []Attribute
	var attrs []JSONAttribute

	err := json.Unmarshal(data, &attrs)
	if err != nil {
		panic(fmt.Errorf("Attribute decode error: %s", err))
	}

	for _, a := range attrs {
		var attr Attribute
		var err error
		switch a.Type {
		case "binary":
			attr = new(BinaryAttribute)
			break
		case "float":
			attr = new(FloatAttribute)
			break
		case "categorical":
			attr = new(CategoricalAttribute)
			break
		default:
			panic(fmt.Errorf("Unrecognised Attribute format: %s", a.Type))
		}
		err = attr.UnmarshalJSON(a.Attr)
		if err != nil {
			panic(fmt.Errorf("Can't deserialize: %s (error: %s)", a, err))
		}
		attr.SetName(a.Name)
		ret = append(ret, attr)
	}
	return ret
}

func DeserializeInstances(f io.Reader) (ret *DenseInstances, err error) {

	// Recovery function
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	// Open the .gz layer
	gzReader, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("Can't open: %s", err)
	}
	// Open the .tar layer
	tr := tar.NewReader(gzReader)
	// Retrieve the MANIFEST and verify
	manifestBytes := getTarContent(tr, "MANIFEST")
	if !reflect.DeepEqual(manifestBytes, []byte(SerializationFormatVersion)) {
		return nil, fmt.Errorf("Unsupported MANIFEST: %s", string(manifestBytes))
	}

	// Get the size
	sizeBytes := getTarContent(tr, "DIMS")
	attrCount := int(UnpackBytesToU64(sizeBytes[0:8]))
	rowCount := int(UnpackBytesToU64(sizeBytes[8:]))

	// Unmarshal the Attributes
	attrBytes := getTarContent(tr, "CATTRS")
	cAttrs := deserializeAttributes(attrBytes)
	attrBytes = getTarContent(tr, "ATTRS")
	normalAttrs := deserializeAttributes(attrBytes)

	// Create the return instances
	ret = NewDenseInstances()

	// Normal Attributes first, class Attributes on the end
	allAttributes := make([]Attribute, attrCount)
	for i, v := range normalAttrs {
		ret.AddAttribute(v)
		allAttributes[i] = v
	}
	for i, v := range cAttrs {
		ret.AddAttribute(v)
		err = ret.AddClassAttribute(v)
		if err != nil {
			return nil, fmt.Errorf("Could not set Attribute as class Attribute: %s", err)
		}
		allAttributes[i+len(normalAttrs)] = v
	}
	// Allocate memory
	err = ret.Extend(int(rowCount))
	if err != nil {
		return nil, fmt.Errorf("Could not allocate memory")
	}

	// Seek through the TAR file until we get to the DATA section
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil, fmt.Errorf("DATA section missing!")
		} else if err != nil {
			return nil, fmt.Errorf("Error seeking to DATA section: %s", err)
		}
		if hdr.Name == "DATA" {
			break
		}
	}

	// Resolve AttributeSpecs
	specs := ResolveAttributes(ret, allAttributes)

	// Finally, read the values out of the data section
	for i := 0; i < rowCount; i++ {
		for _, s := range specs {
			r := ret.Get(s, i)
			n, err := tr.Read(r)
			if n != len(r) {
				return nil, fmt.Errorf("Expected %d bytes (read %d) on row %d", len(r), n, i)
			}
			if err != nil {
				return nil, fmt.Errorf("Read error: %s", err)
			}
			ret.Set(s, i, r)
		}
	}

	if err = gzReader.Close(); err != nil {
		return ret, fmt.Errorf("Error closing gzip stream: %s", err)
	}

	return ret, nil
}

func SerializeInstances(inst FixedDataGrid, f io.Writer) error {
	var hdr *tar.Header

	gzWriter := gzip.NewWriter(f)
	tw := tar.NewWriter(gzWriter)

	// Write the MANIFEST entry
	hdr = &tar.Header{
		Name: "MANIFEST",
		Size: int64(len(SerializationFormatVersion)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Could not write MANIFEST header: %s", err)
	}

	if _, err := tw.Write([]byte(SerializationFormatVersion)); err != nil {
		return fmt.Errorf("Could not write MANIFEST contents: %s", err)
	}

	// Now write the dimensions of the dataset
	attrCount, rowCount := inst.Size()
	hdr = &tar.Header{
		Name: "DIMS",
		Size: 16,
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Could not write DIMS header: %s", err)
	}

	if _, err := tw.Write(PackU64ToBytes(uint64(attrCount))); err != nil {
		return fmt.Errorf("Could not write DIMS (attrCount): %s", err)
	}
	if _, err := tw.Write(PackU64ToBytes(uint64(rowCount))); err != nil {
		return fmt.Errorf("Could not write DIMS (rowCount): %s", err)
	}

	// Write the ATTRIBUTES files
	classAttrs := inst.AllClassAttributes()
	normalAttrs := NonClassAttributes(inst)
	if err := writeAttributesToFilePart(classAttrs, tw, "CATTRS"); err != nil {
		return fmt.Errorf("Could not write CATTRS: %s", err)
	}
	if err := writeAttributesToFilePart(normalAttrs, tw, "ATTRS"); err != nil {
		return fmt.Errorf("Could not write ATTRS: %s", err)
	}

	// Data must be written out in the same order as the Attributes
	allAttrs := make([]Attribute, attrCount)
	normCount := copy(allAttrs, normalAttrs)
	for i, v := range classAttrs {
		allAttrs[normCount+i] = v
	}

	allSpecs := ResolveAttributes(inst, allAttrs)

	// First, estimate the amount of data we'll need...
	dataLength := int64(0)
	inst.MapOverRows(allSpecs, func(val [][]byte, row int) (bool, error) {
		for _, v := range val {
			dataLength += int64(len(v))
		}
		return true, nil
	})

	// Then write the header
	hdr = &tar.Header{
		Name: "DATA",
		Size: dataLength,
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Could not write DATA: %s", err)
	}

	// Then write the actual data
	writtenLength := int64(0)
	if err := inst.MapOverRows(allSpecs, func(val [][]byte, row int) (bool, error) {
		for _, v := range val {
			wl, err := tw.Write(v)
			writtenLength += int64(wl)
			if err != nil {
				return false, err
			}
		}
		return true, nil
	}); err != nil {
		return err
	}

	if writtenLength != dataLength {
		return fmt.Errorf("Could not write DATA: changed size from %v to %v", dataLength, writtenLength)
	}

	// Finally, close and flush the various levels
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("Could not flush tar: %s", err)
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("Could not close tar: %s", err)
	}

	if err := gzWriter.Flush(); err != nil {
		return fmt.Errorf("Could not flush gz: %s", err)
	}

	if err := gzWriter.Close(); err != nil {
		return fmt.Errorf("Could not close gz: %s", err)
	}

	return nil
}
