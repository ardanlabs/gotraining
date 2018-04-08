package base

import (
	"archive/tar"
	"encoding/json"
	"fmt"
)

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

// MarshalAttribute converts an Attribute to a JSON map.
func MarshalAttribute(a Attribute) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	marshaledAttrRaw, err := a.MarshalJSON()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshaledAttrRaw, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func SerializeAttribute(attr Attribute) ([]byte, error) {
	// Get the marshaled Attribute array
	body, err := json.Marshal(attr)
	if err != nil {
		return nil, err
	}
	return []byte(body), nil
}

func DeserializeAttribute(data []byte) (Attribute, error) {
	type JSONAttribute struct {
		Type string          `json:"type"`
		Name string          `json:"name"`
		Attr json.RawMessage `json:"attr"`
	}

	var rawAttr JSONAttribute
	err := json.Unmarshal(data, &rawAttr)
	if err != nil {
		return nil, err
	}
	var attr Attribute

	switch rawAttr.Type {
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
		return nil, fmt.Errorf("Unrecognised Attribute format: %s", rawAttr.Type)
	}

	err = attr.UnmarshalJSON(rawAttr.Attr)
	if err != nil {
		return nil, fmt.Errorf("Can't deserialize: %s (error: %s)", rawAttr, err)
	}
	attr.SetName(rawAttr.Name)
	return attr, nil
}

// DeserializeAttributes constructs a ve
func DeserializeAttributes(data []byte) ([]Attribute, error) {

	// Define a JSON shim Attribute
	var attrs []json.RawMessage
	err := json.Unmarshal(data, &attrs)
	if err != nil {
		return nil, fmt.Errorf("Failed to deserialize attributes: %v", err)
	}

	ret := make([]Attribute, len(attrs))
	for i, v := range attrs {
		ret[i], err = DeserializeAttribute(v)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ReplaceDeserializedAttributeWithVersionFromInstances takes an independently deserialized Attribute and matches it
// if possible with one from a candidate FixedDataGrid.
func ReplaceDeserializedAttributeWithVersionFromInstances(deserialized Attribute, matchingWith FixedDataGrid) (Attribute, error) {
	for _, a := range matchingWith.AllAttributes() {
		if a.Equals(deserialized) {
			return a, nil
		}
	}
	return nil, WrapError(fmt.Errorf("Unable to match %v in %v", deserialized, matchingWith))
}

// ReplaceDeserializedAttributesWithVersionsFromInstances takes some independently loaded Attributes and
// matches them up with a candidate FixedDataGrid.
func ReplaceDeserializedAttributesWithVersionsFromInstances(deserialized []Attribute, matchingWith FixedDataGrid) ([]Attribute, error) {
	ret := make([]Attribute, len(deserialized))
	for i, a := range deserialized {
		match, err := ReplaceDeserializedAttributeWithVersionFromInstances(a, matchingWith)
		if err != nil {
			return nil, WrapError(err)
		}
		ret[i] = match
	}
	return ret, nil
}
