// Copyright (C) 2011, Ross Light

package pdf

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// A marshaler can produce a PDF object.
type marshaler interface {
	marshalPDF(dst []byte) ([]byte, error)
}

// marshal returns the PDF encoding of v.
//
// If the value implements the marshaler interface, then its marshalPDF method
// is called.  ints, strings, and floats will be marshalled according to the PDF
// standard.
func marshal(dst []byte, v interface{}) ([]byte, error) {
	state := marshalState{dst}
	if err := state.marshalValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return state.data, nil
}

type marshalState struct {
	data []byte
}

const marshalFloatPrec = 5

func (state *marshalState) writeString(s string) {
	state.data = append(state.data, s...)
}

func (state *marshalState) marshalValue(v reflect.Value) error {
	if !v.IsValid() {
		state.writeString("null")
		return nil
	}

	if m, ok := v.Interface().(marshaler); ok {
		slice, err := m.marshalPDF(state.data)
		if err != nil {
			return err
		}
		state.data = slice
		return nil
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		state.writeString(strconv.FormatInt(v.Int(), 10))
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		state.writeString(strconv.FormatUint(v.Uint(), 10))
		return nil
	case reflect.Float32, reflect.Float64:
		state.writeString(strconv.FormatFloat(v.Float(), 'f', marshalFloatPrec, 64))
		return nil
	case reflect.String:
		state.writeString(quote(v.String()))
		return nil
	case reflect.Ptr, reflect.Interface:
		return state.marshalValue(v.Elem())
	case reflect.Array, reflect.Slice:
		return state.marshalSlice(v)
	case reflect.Map:
		return state.marshalDictionary(v)
	case reflect.Struct:
		return state.marshalStruct(v)
	}

	return errors.New("pdf: unsupported type: " + v.Type().String())
}

// quote escapes a string and returns a PDF string literal.
func quote(s string) string {
	r := strings.NewReplacer(
		"\r", `\r`,
		"\t", `\t`,
		"\b", `\b`,
		"\f", `\f`,
		"(", `\(`,
		")", `\)`,
		`\`, `\\`,
	)
	return "(" + r.Replace(s) + ")"
}

func (state *marshalState) marshalSlice(v reflect.Value) error {
	state.writeString("[ ")
	for i := 0; i < v.Len(); i++ {
		if err := state.marshalValue(v.Index(i)); err != nil {
			return err
		}
		state.writeString(" ")
	}
	state.writeString("]")
	return nil
}

func (state *marshalState) marshalDictionary(v reflect.Value) error {
	if v.Type().Key() != reflect.TypeOf(name("")) {
		return errors.New("pdf: cannot marshal dictionary with key type: " + v.Type().Key().String())
	}

	state.writeString("<< ")
	for _, k := range v.MapKeys() {
		state.marshalKeyValue(k.Interface().(name), v.MapIndex(k))
	}
	state.writeString(">>")
	return nil
}

func (state *marshalState) marshalStruct(v reflect.Value) error {
	state.writeString("<< ")
	t := v.Type()
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}

		tag, omitEmpty := f.Name, false
		if tv := f.Tag.Get("pdf"); tv != "" {
			if tv == "-" {
				continue
			}

			name, options := parseTag(tv)
			if name != "" {
				tag = name
			}
			omitEmpty = options.Contains("omitempty")
		}

		fieldValue := v.Field(i)
		if omitEmpty && isEmptyValue(fieldValue) {
			continue
		}

		state.marshalKeyValue(name(tag), fieldValue)
	}
	state.writeString(">>")
	return nil
}

func (state *marshalState) marshalKeyValue(k name, v reflect.Value) error {
	slice, err := k.marshalPDF(state.data)
	if err != nil {
		return err
	}
	state.data = slice
	state.writeString(" ")

	if err := state.marshalValue(v); err != nil {
		return err
	}
	state.writeString(" ")

	return nil
}

type tagOptions []string

func parseTag(tag string) (name string, options tagOptions) {
	result := strings.Split(tag, ",")
	return result[0], tagOptions(result[1:])
}

func (options tagOptions) Contains(opt string) bool {
	for _, o := range options {
		if opt == o {
			return true
		}
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
