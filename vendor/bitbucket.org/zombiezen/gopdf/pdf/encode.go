// Copyright (C) 2011, Ross Light

package pdf

import (
	"fmt"
	"io"
)

// encoder writes the PDF file format structure.
type encoder struct {
	objects []interface{}
	root    Reference
}

type trailer struct {
	Size int
	Root Reference
}

// add appends an object to the file.  The object is marshalled only when an
// encoding is requested.
func (enc *encoder) add(v interface{}) Reference {
	enc.objects = append(enc.objects, v)
	return Reference{uint(len(enc.objects)), 0}
}

const (
	header  = "%PDF-1.7" + newline + "%\x93\x8c\x8b\x9e" + newline
	newline = "\r\n"
)

// Cross reference strings
const (
	crossReferenceSectionHeader    = "xref" + newline
	crossReferenceSubsectionFormat = "%d %d" + newline
	crossReferenceFormat           = "%010d %05d n" + newline
	crossReferenceFreeFormat       = "%010d %05d f" + newline
)

const trailerHeader = "trailer" + newline

const startxrefFormat = "startxref" + newline + "%d" + newline

const eofString = "%%EOF" + newline

// encode writes an entire PDF document by marshalling the added objects.
func (enc *encoder) encode(wr io.Writer) error {
	w := &offsetWriter{Writer: wr}
	if err := enc.writeHeader(w); err != nil {
		return err
	}
	objectOffsets, err := enc.writeBody(w)
	if err != nil {
		return err
	}
	tableOffset := w.offset
	if err := enc.writeXrefTable(w, objectOffsets); err != nil {
		return err
	}
	if err := enc.writeTrailer(w); err != nil {
		return err
	}
	if err := enc.writeStartxref(w, tableOffset); err != nil {
		return err
	}
	if err := enc.writeEOF(w); err != nil {
		return err
	}
	return nil
}

func (enc *encoder) writeHeader(w *offsetWriter) error {
	_, err := io.WriteString(w, header)
	return err
}

func (enc *encoder) writeBody(w *offsetWriter) ([]int64, error) {
	objectOffsets := make([]int64, len(enc.objects))
	for i, obj := range enc.objects {
		// TODO: Use same buffer for writing across objects
		objectOffsets[i] = w.offset
		data, err := marshal(nil, indirectObject{Reference{uint(i + 1), 0}, obj})
		if err != nil {
			return nil, err
		}
		if _, err = w.Write(append(data, newline...)); err != nil {
			return nil, err
		}
	}
	return objectOffsets, nil
}

func (enc *encoder) writeXrefTable(w *offsetWriter, objectOffsets []int64) error {
	if _, err := io.WriteString(w, crossReferenceSectionHeader); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, crossReferenceSubsectionFormat, 0, len(enc.objects)+1); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, crossReferenceFreeFormat, 0, 65535); err != nil {
		return err
	}
	for _, offset := range objectOffsets {
		if _, err := fmt.Fprintf(w, crossReferenceFormat, offset, 0); err != nil {
			return err
		}
	}
	return nil
}

func (enc *encoder) writeTrailer(w *offsetWriter) error {
	var err error
	dict := trailer{
		Size: len(enc.objects) + 1,
		Root: enc.root,
	}
	data := make([]byte, 0, len(trailerHeader)+len(newline))
	data = append(data, trailerHeader...)
	if data, err = marshal(data, dict); err != nil {
		return err
	}
	data = append(data, newline...)

	_, err = w.Write(data)
	return err
}

func (enc *encoder) writeStartxref(w *offsetWriter, tableOffset int64) error {
	_, err := fmt.Fprintf(w, startxrefFormat, tableOffset)
	return err
}

func (enc *encoder) writeEOF(w *offsetWriter) error {
	_, err := io.WriteString(w, eofString)
	return err
}

// offsetWriter tracks how many bytes have been written to it.
type offsetWriter struct {
	io.Writer
	offset int64
}

func (w *offsetWriter) Write(p []byte) (n int, err error) {
	n, err = w.Writer.Write(p)
	w.offset += int64(n)
	return
}
