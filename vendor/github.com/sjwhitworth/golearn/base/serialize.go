package base

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

const (
	SerializationFormatVersion = "golearn 1.0"
)

// FunctionalTarReader allows you to read anything in a tar file in any order, rather than just
// sequentially.
type FunctionalTarReader struct {
	Regenerate func() *tar.Reader
}

// NewFunctionalTarReader creates a new FunctionalTarReader using a function that it can call
// to get a tar.Reader at the beginning of the file.
func NewFunctionalTarReader(regenFunc func() *tar.Reader) *FunctionalTarReader {
	return &FunctionalTarReader{
		regenFunc,
	}
}

// GetNamedFile returns a file named a given thing from the tar file. If there's more than one
// entry, the most recent is returned.
func (f *FunctionalTarReader) GetNamedFile(name string) ([]byte, error) {
	tr := f.Regenerate()

	var returnCandidate []byte = nil
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if hdr.Name == name {
			ret, err := ioutil.ReadAll(tr)
			if err != nil {
				return nil, WrapError(err)
			}
			if int64(len(ret)) != hdr.Size {
				if int64(len(ret)) < hdr.Size {
					log.Printf("Size mismatch, got %d byte(s) for %s, expected %d (err was %s)", len(ret), hdr.Name, hdr.Size, err)
				} else {
					return nil, WrapError(fmt.Errorf("Size mismatch, expected %d byte(s) for %s, got %d", len(ret), hdr.Name, hdr.Size))
				}
			}
			if err != nil {
				return nil, err
			}
			returnCandidate = ret
		}
	}
	if returnCandidate == nil {
		return nil, WrapError(fmt.Errorf("Not found (looking for %s)", name))
	}
	return returnCandidate, nil
}

func tarPrefix(prefix string, suffix string) string {
	if prefix == "" {
		return suffix
	}
	return fmt.Sprintf("%s/%s", prefix, suffix)
}

// ClassifierMetadataV1 is what gets written into METADATA
// in a classification file format.
type ClassifierMetadataV1 struct {
	// FormatVersion should always be 1 for this structure
	FormatVersion int `json:"format_version"`
	// Uses the classifier name (provided by the classifier)
	ClassifierName string `json:"classifier"`
	// ClassifierVersion is also provided by the classifier
	// and checks whether this version of GoLearn can read what's
	// be written.
	ClassifierVersion string `json"classifier_version"`
	// This is a custom metadata field, provided by the classifier
	ClassifierMetadata map[string]interface{} `json:"classifier_metadata"`
}

// ClassifierDeserializer attaches helper functions useful for reading classificatiers. (UNSTABLE).
type ClassifierDeserializer struct {
	gzipReader io.Reader
	fileReader io.ReadCloser
	tarReader  *FunctionalTarReader
	Metadata   *ClassifierMetadataV1
}

// Prefix outputs a string in the right format for TAR
func (c *ClassifierDeserializer) Prefix(prefix string, suffix string) string {
	if prefix == "" {
		return suffix
	}
	return fmt.Sprintf("%s/%s", prefix, suffix)
}

// ReadMetadataAtPrefix reads the METADATA file after prefix. If an error is returned, the first value is undefined.
func (c *ClassifierDeserializer) ReadMetadataAtPrefix(prefix string) (ClassifierMetadataV1, error) {
	var ret ClassifierMetadataV1
	err := c.GetJSONForKey(c.Prefix(prefix, "METADATA"), &ret)
	return ret, err
}

// ReadSerializedClassifierStub is the counterpart of CreateSerializedClassifierStub.
// It's used inside SaveableClassifiers to read information from a perviously saved
// model file.
func ReadSerializedClassifierStub(filePath string) (*ClassifierDeserializer, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, DescribeError("Can't open file", err)
	}

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return nil, DescribeError("Can't decompress", err)
	}

	regenerateFunc := func() *tar.Reader {
		f.Seek(0, os.SEEK_SET)
		gzr.Reset(f)
		tz := tar.NewReader(gzr)
		return tz
	}

	tz := NewFunctionalTarReader(regenerateFunc)

	// Check that the serialization format is right
	// Retrieve the MANIFEST and verify
	manifestBytes, err := tz.GetNamedFile("CLS_MANIFEST")
	if err != nil {
		return nil, DescribeError("Error reading CLS_MANIFEST", err)
	}
	if !reflect.DeepEqual(manifestBytes, []byte(SerializationFormatVersion)) {
		return nil, fmt.Errorf("Unsupported CLS_MANIFEST: %s", string(manifestBytes))
	}

	//
	// Parse METADATA
	//
	var metadata ClassifierMetadataV1
	ret := &ClassifierDeserializer{
		f,
		gzr,
		tz,
		&metadata,
	}

	metadata, err = ret.ReadMetadataAtPrefix("")
	if err != nil {
		return nil, fmt.Errorf("Error whilst reading METADATA: %s", err)
	}
	ret.Metadata = &metadata

	// Check that we can understand this archive
	if metadata.FormatVersion != 1 {
		return nil, fmt.Errorf("METADATA: wrong format_version for this version of golearn")
	}

	return ret, nil
}

// GetBytesForKey returns the bytes at a given location in the output.
func (c *ClassifierDeserializer) GetBytesForKey(key string) ([]byte, error) {
	return c.tarReader.GetNamedFile(key)
}

func (c *ClassifierDeserializer) GetStringForKey(key string) (string, error) {
	b, err := c.GetBytesForKey(key)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// GetJSONForKey deserializes a JSON key in the output file.
func (c *ClassifierDeserializer) GetJSONForKey(key string, v interface{}) error {
	b, err := c.GetBytesForKey(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// GetInstancesForKey deserializes some instances stored in a classifier output file
func (c *ClassifierDeserializer) GetInstancesForKey(key string) (FixedDataGrid, error) {
	return DeserializeInstancesFromTarReader(c.tarReader, key)
}

// GetUInt64ForKey returns a int64 stored at a given key
func (c *ClassifierDeserializer) GetU64ForKey(key string) (uint64, error) {
	b, err := c.GetBytesForKey(key)
	if err != nil {
		return 0, err
	}
	return UnpackBytesToU64(b), nil
}

// GetAttributeForKey returns an Attribute stored at a given key
func (c *ClassifierDeserializer) GetAttributeForKey(key string) (Attribute, error) {
	b, err := c.GetBytesForKey(key)
	if err != nil {
		return nil, WrapError(err)
	}
	attr, err := DeserializeAttribute(b)
	if err != nil {
		return nil, WrapError(err)
	}
	return attr, nil
}

// GetAttributesForKey returns an Attribute list stored at a given key
func (c *ClassifierDeserializer) GetAttributesForKey(key string) ([]Attribute, error) {

	attrCountKey := c.Prefix(key, "ATTR_COUNT")
	attrCount, err := c.GetU64ForKey(attrCountKey)
	if err != nil {
		return nil, DescribeError("Unable to read ATTR_COUNT", err)
	}

	ret := make([]Attribute, attrCount)

	for i := range ret {
		attrKey := c.Prefix(key, fmt.Sprintf("%d", i))
		ret[i], err = c.GetAttributeForKey(attrKey)
		if err != nil {
			return nil, DescribeError("Unable to read Attribute", err)
		}
	}
	return ret, nil
}

// Close cleans up everything.
func (c *ClassifierDeserializer) Close() {
	c.fileReader.Close()
}

// ClassifierSerializer is an object used by SaveableClassifiers.
type ClassifierSerializer struct {
	gzipWriter *gzip.Writer
	fileWriter *os.File
	tarWriter  *tar.Writer
	f          *os.File
	filePath   string
}

// Close finalizes the Classifier serialization session.
func (c *ClassifierSerializer) Close() error {

	// Finally, close and flush the various levels
	if err := c.tarWriter.Flush(); err != nil {
		return fmt.Errorf("Could not flush tar: %s", err)
	}

	if err := c.tarWriter.Close(); err != nil {
		return fmt.Errorf("Could not close tar: %s", err)
	}

	if err := c.gzipWriter.Flush(); err != nil {
		return fmt.Errorf("Could not flush gz: %s", err)
	}

	if err := c.gzipWriter.Close(); err != nil {
		return fmt.Errorf("Could not close gz: %s", err)
	}

	if err := c.fileWriter.Sync(); err != nil {
		return fmt.Errorf("Could not close file writer: %s", err)
	}

	if err := c.fileWriter.Close(); err != nil {
		return fmt.Errorf("Could not close file writer: %s", err)
	}

	return nil
}

// WriteBytesForKey creates a new entry in the serializer file with some user-defined bytes.
func (c *ClassifierSerializer) WriteBytesForKey(key string, b []byte) error {

	//
	// Write header for key
	//
	hdr := &tar.Header{
		Name: key,
		Size: int64(len(b)),
	}

	if err := c.tarWriter.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Could not write header for '%s': %s", key, err)
	}
	//
	// Write data
	//
	if _, err := c.tarWriter.Write(b); err != nil {
		return fmt.Errorf("Could not write data for '%s': %s", key, err)
	}

	c.tarWriter.Flush()
	c.gzipWriter.Flush()
	c.fileWriter.Sync()
	return nil
}

// WriteU64ForKey creates a new entry in the serializer file with the bytes of a uint64
func (c *ClassifierSerializer) WriteU64ForKey(key string, v uint64) error {
	b := PackU64ToBytes(v)
	return c.WriteBytesForKey(key, b)
}

// WriteJSONForKey creates a new entry in the file with an interface serialized as JSON.
func (c *ClassifierSerializer) WriteJSONForKey(key string, v interface{}) error {

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return c.WriteBytesForKey(key, b)

}

// WriteAttributeForKey creates a new entry in the file containing a serialized representation of Attribute
func (c *ClassifierSerializer) WriteAttributeForKey(key string, a Attribute) error {
	b, err := SerializeAttribute(a)
	if err != nil {
		return WrapError(err)
	}
	return c.WriteBytesForKey(key, b)
}

// WriteAttributesForKey does the same as WriteAttributeForKey, just with more than one Attribute.
func (c *ClassifierSerializer) WriteAttributesForKey(key string, attrs []Attribute) error {

	attrCountKey := c.Prefix(key, "ATTR_COUNT")
	err := c.WriteU64ForKey(attrCountKey, uint64(len(attrs)))
	if err != nil {
		return DescribeError("Unable to write ATTR_COUNT", err)
	}
	for i, a := range attrs {
		attrKey := c.Prefix(key, fmt.Sprintf("%d", i))
		err = c.WriteAttributeForKey(attrKey, a)
		if err != nil {
			return DescribeError("Unable to write Attribute", err)
		}
	}
	return nil
}

// WriteInstances for key creates a new entry in the file containing some training instances
func (c *ClassifierSerializer) WriteInstancesForKey(key string, g FixedDataGrid, includeData bool) error {
	fmt.Sprintf("%v", c)
	return SerializeInstancesToTarWriter(g, c.tarWriter, key, includeData)
}

// Prefix outputs a string in the right format for TAR
func (c *ClassifierSerializer) Prefix(prefix string, suffix string) string {
	if prefix == "" {
		return suffix
	}
	return fmt.Sprintf("%s/%s", prefix, suffix)
}

// WriteMetadataAtPrefix outputs a METADATA entry in the right place
func (c *ClassifierSerializer) WriteMetadataAtPrefix(prefix string, metadata ClassifierMetadataV1) error {
	return c.WriteJSONForKey(c.Prefix(prefix, "METADATA"), &metadata)
}

// CreateSerializedClassifierStub generates a file to serialize into
// and writes the METADATA header.
func CreateSerializedClassifierStub(filePath string, metadata ClassifierMetadataV1) (*ClassifierSerializer, error) {

	// Open the filePath
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	var hdr *tar.Header
	gzWriter := gzip.NewWriter(f)
	tw := tar.NewWriter(gzWriter)

	ret := ClassifierSerializer{
		gzipWriter: gzWriter,
		fileWriter: f,
		tarWriter:  tw,
	}

	//
	// Write the MANIFEST entry
	//
	hdr = &tar.Header{
		Name: "CLS_MANIFEST",
		Size: int64(len(SerializationFormatVersion)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return nil, fmt.Errorf("Could not write CLS_MANIFEST header: %s", err)
	}

	if _, err := tw.Write([]byte(SerializationFormatVersion)); err != nil {
		return nil, fmt.Errorf("Could not write CLS_MANIFEST contents: %s", err)
	}

	//
	// Write the METADATA entry
	//
	err = ret.WriteMetadataAtPrefix("", metadata)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %s", err)
	}

	return &ret, nil

}
