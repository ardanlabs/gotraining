// Copyright (C) 2011, Ross Light

package pdf

import (
	"fmt"
	"strconv"
)

// name is a PDF name object, which is used as an identifier.
type name string

func (n name) String() string {
	return string(n)
}

func (n name) marshalPDF(dst []byte) ([]byte, error) {
	// TODO: escape characters
	dst = append(dst, '/')
	return append(dst, []byte(n)...), nil
}

type indirectObject struct {
	Reference
	Object interface{}
}

const (
	objectBegin = " obj\r\n"
	objectEnd   = "\r\nendobj"
)

func (obj indirectObject) marshalPDF(dst []byte) ([]byte, error) {
	var err error
	mn, mg := strconv.FormatUint(uint64(obj.Number), 10), strconv.FormatUint(uint64(obj.Generation), 10)
	dst = append(dst, mn...)
	dst = append(dst, ' ')
	dst = append(dst, mg...)
	dst = append(dst, objectBegin...)
	if dst, err = marshal(dst, obj.Object); err != nil {
		return nil, err
	}
	dst = append(dst, objectEnd...)
	return dst, nil
}

// Reference holds a PDF indirect reference.
type Reference struct {
	Number     uint
	Generation uint
}

func (ref Reference) marshalPDF(dst []byte) ([]byte, error) {
	return append(dst, fmt.Sprintf("%d %d R", ref.Number, ref.Generation)...), nil
}
