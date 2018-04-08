package pbutil

import (
	"encoding/binary"
	"io"

	"github.com/gogo/protobuf/proto"
)

// Reader is io.Reader for proto.Message instead of []byte.
type Reader interface {
	Read(val proto.Message) error
}

// Writer is io.Writer for proto.Message instead of []byte.
type Writer interface {
	Write(val proto.Message) error
}

// ReadWriter is io.ReadWriter for proto.Message instead of []byte.
type ReadWriter interface {
	Reader
	Writer
}

type readWriter struct {
	w   io.Writer
	r   io.Reader
	buf []byte
}

// Read reads val from r.
func (r *readWriter) Read(val proto.Message) error {
	var l int64
	if err := binary.Read(r.r, binary.LittleEndian, &l); err != nil {
		return err
	}
	if r.buf == nil || len(r.buf) < int(l) {
		r.buf = make([]byte, l)
	}
	buf := r.buf[0:l]
	if _, err := io.ReadFull(r.r, buf); err != nil {
		if err == io.EOF {
			return io.ErrUnexpectedEOF
		}
		return err
	}
	return proto.Unmarshal(buf, val)
}

// Write writes val to r.
func (r *readWriter) Write(val proto.Message) error {
	bytes, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	if err := binary.Write(r.w, binary.LittleEndian, int64(len(bytes))); err != nil {
		return err
	}
	_, err = r.w.Write(bytes)
	return err
}

// NewReader returns a new Reader with r as its source.
func NewReader(r io.Reader) Reader {
	return &readWriter{r: r}
}

// NewWriter returns a new Writer with w as its sink.
func NewWriter(w io.Writer) Writer {
	return &readWriter{w: w}
}

// NewReadWriter returns a new ReadWriter with rw as both its source and its sink.
func NewReadWriter(rw io.ReadWriter) ReadWriter {
	return &readWriter{r: rw, w: rw}
}
