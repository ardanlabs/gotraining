package base

import (
	"archive/tar"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
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

// SerializesInstancesToCSV converts a FixedDataGrid into a CSV file format.
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

// SerializeInstancesToCSVStream outputs a FixedDataGrid into a CSV file format, via the io.Writer stream.
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

// DeserializeInstancesFromTarReader returns DenseInstances from a FunctionalTarReader with the name prefix.
func DeserializeInstancesFromTarReader(tr *FunctionalTarReader, prefix string) (ret *DenseInstances, err error) {

	p := func(n string) string {
		return fmt.Sprintf("%s%s", prefix, n)
	}

	// Retrieve the MANIFEST and verify
	manifestBytes, err := tr.GetNamedFile(p("MANIFEST"))
	if err != nil {
		return nil, err
	}
	if !reflect.DeepEqual(manifestBytes, []byte(SerializationFormatVersion)) {
		return nil, fmt.Errorf("Unsupported MANIFEST: %s", string(manifestBytes))
	}

	// Get the size
	sizeBytes, err := tr.GetNamedFile(p("DIMS"))
	if err != nil {
		return nil, WrapError(fmt.Errorf("Unable to read DIMS: %v", err))
	}
	if len(sizeBytes) < 16 {
		return nil, WrapError(fmt.Errorf("DIMS: must be 16 bytes"))
	}
	attrCount := int(UnpackBytesToU64(sizeBytes[0:8]))
	rowCount := int(UnpackBytesToU64(sizeBytes[8:]))

	// Unmarshal the Attributes
	attrBytes, err := tr.GetNamedFile(p("CATTRS"))
	if err != nil {
		return nil, DescribeError("Unable to read CATTRS", err)
	}
	cAttrs, err := DeserializeAttributes(attrBytes)
	if err != nil {
		return nil, DescribeError("Class Attribute deserialization error", err)
	}
	attrBytes, err = tr.GetNamedFile(p("ATTRS"))
	if err != nil {
		return nil, DescribeError("Unable to read ATTRS", err)
	}
	normalAttrs, err := DeserializeAttributes(attrBytes)
	if err != nil {
		return nil, DescribeError("Unable to deserialize normal attributes", err)
	}

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
			return nil, DescribeError(fmt.Sprintf("Could not set Attribute '%s' as a class Attribute", v), err)
		}
		allAttributes[i+len(normalAttrs)] = v
	}
	// Allocate memory
	err = ret.Extend(int(rowCount))
	if err != nil {
		return nil, WrapError(fmt.Errorf("Could not allocate memory"))
	}

	// Seek through the TAR file until we get to the DATA section
	reader := tr.Regenerate()
	for {
		hdr, err := reader.Next()
		if err == io.EOF {
			return nil, WrapError(fmt.Errorf("DATA section missing!"))
		} else if err != nil {
			return nil, WrapError(fmt.Errorf("Error seeking to DATA section: %s", err))
		}
		if hdr.Name == p("DATA") {
			break
		}
	}

	// Resolve AttributeSpecs
	specs := ResolveAttributes(ret, allAttributes)

	// Finally, read the values out of the data section
	for i := 0; i < rowCount; i++ {
		for j, s := range specs {
			r := ret.Get(s, i)
			n, err := reader.Read(r)
			if n != len(r) {
				return nil, WrapError(fmt.Errorf("Expected %d bytes (read %d) on row %d", len(r), n, i))
			}
			ret.Set(s, i, r)
			if err != nil {
				if i == rowCount-1 && j == len(specs)-1 && err == io.EOF {
					break
				}
				return nil, WrapError(fmt.Errorf("Read error in data section (at row %d from %d, attr %d from %d): %s", i, rowCount, j, len(specs), err))
			}
		}
	}

	return ret, nil
}

// DeserializeInstances returns a DenseInstances using a given io.Reader.
func DeserializeInstances(f io.ReadSeeker) (ret *DenseInstances, err error) {

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
		panic(WrapError(err))
	}
	regenerateTarReader := func() *tar.Reader {
		f.Seek(0, os.SEEK_SET)
		gzReader.Reset(f)
		tr := tar.NewReader(gzReader)
		return tr
	}

	tr := NewFunctionalTarReader(regenerateTarReader)

	ret, deSerializeErr := DeserializeInstancesFromTarReader(tr, "")

	if err = gzReader.Close(); err != nil {
		return ret, fmt.Errorf("Error closing gzip stream: %s", err)
	}

	return ret, deSerializeErr
}

// SerializeInstances stores a FixedDataGrid into an efficient format to the given io.Writer stream.
func SerializeInstances(inst FixedDataGrid, f io.Writer) error {
	// Create a .tar.gz container
	gzWriter := gzip.NewWriter(f)
	tw := tar.NewWriter(gzWriter)

	serializeErr := SerializeInstancesToTarWriter(inst, tw, "", true)
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

	return serializeErr
}

// SerializeInstancesToTarWriter stores a FixedDataGrid into an efficient form given a tar.Writer.
func SerializeInstancesToTarWriter(inst FixedDataGrid, tw *tar.Writer, prefix string, includeData bool) error {
	var hdr *tar.Header

	p := func(n string) string {
		return fmt.Sprintf("%s%s", prefix, n)
	}

	// Write the MANIFEST entry
	hdr = &tar.Header{
		Name: p("MANIFEST"),
		Size: int64(len(SerializationFormatVersion)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Could not write MANIFEST header: %s", err)
	}

	if _, err := tw.Write([]byte(SerializationFormatVersion)); err != nil {
		return fmt.Errorf("Could not write MANIFEST contents: %s", err)
	}
	tw.Flush()

	// Now write the dimensions of the dataset
	attrCount, rowCount := inst.Size()
	hdr = &tar.Header{
		Name: p("DIMS"),
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
	if err := writeAttributesToFilePart(classAttrs, tw, p("CATTRS")); err != nil {
		return fmt.Errorf("Could not write CATTRS: %s", err)
	}
	if err := writeAttributesToFilePart(normalAttrs, tw, p("ATTRS")); err != nil {
		return fmt.Errorf("Could not write ATTRS: %s", err)
	}

	// Data must be written out in the same order as the Attributes
	allAttrs := make([]Attribute, attrCount)
	normCount := copy(allAttrs, normalAttrs)
	for i, v := range classAttrs {
		allAttrs[normCount+i] = v
	}

	allSpecs := ResolveAttributes(inst, allAttrs)
	if len(allSpecs) != len(allAttrs) {
		return WrapError(fmt.Errorf("Error resolving all Attributes: resolved %d, expected %d", len(allSpecs), len(allAttrs)))
	}

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
		Name: p("DATA"),
		Size: dataLength,
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Could not write DATA: %s", err)
	}
	tw.Flush()

	if !includeData {
		return nil
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

	tw.Flush()

	return nil
}
