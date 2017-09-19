// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package csvutil exposes functions and types to easily handle CSV files.
// csvutil builds upon the encoding/csv package and provides the Table type.
// A Table can read data from a CSV file into a struct value whose fields are
// the various columns of the CSV file.
// Conversely, a Table can write data into a CSV file from a struct value.
package csvutil // import "go-hep.org/x/hep/csvutil"

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Open opens a Table in read mode connected to a CSV file.
func Open(fname string) (*Table, error) {
	r, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	table := &Table{
		Reader: csv.NewReader(bufio.NewReader(r)),
		f:      r,
	}
	return table, err
}

// Create creates a new CSV file and returns a Table in write mode.
func Create(fname string) (*Table, error) {
	w, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	table := &Table{
		Writer: csv.NewWriter(bufio.NewWriter(w)),
		f:      w,
	}
	return table, err
}

// Append opens an already existing CSV file and returns a Table in write mode.
// The file cursor is positioned at the end of the file so new data can be
// appended via the returned Table.
func Append(fname string) (*Table, error) {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, os.SEEK_END)
	if err != nil {
		return nil, err
	}

	table := &Table{
		Writer: csv.NewWriter(bufio.NewWriter(f)),
		f:      f,
	}
	return table, err
}

// Table provides read- or write-access to a CSV file.
// Table supports reading and writing data to/from a struct value.
type Table struct {
	Reader *csv.Reader
	Writer *csv.Writer

	f      *os.File
	closed bool
	err    error
}

// Close closes the table and the underlying CSV file.
func (tbl *Table) Close() error {
	if tbl.closed {
		return tbl.err
	}

	if tbl.Writer != nil {
		tbl.Writer.Flush()
		tbl.err = tbl.Writer.Error()
	}

	if tbl.f != nil {
		err := tbl.f.Close()
		if err != nil && tbl.err == nil {
			tbl.err = err
		}
		tbl.f = nil
		tbl.closed = true
	}
	return tbl.err
}

// ReadRows returns a row iterator semantically equivalent to [beg,end).
// If end==-1, the iterator will be configured to read rows until EOF.
func (tbl *Table) ReadRows(beg, end int64) (*Rows, error) {
	inc := int64(1)
	rows := &Rows{
		tbl: tbl,
		i:   0,
		n:   end - beg,
		inc: inc,
		cur: beg - inc,
	}
	if end == -1 {
		rows.n = math.MaxInt64
	}
	if beg > 0 {
		err := rows.skip(beg)
		if err != nil {
			return nil, err
		}
	}
	return rows, nil
}

// WriteHeader writes a header to the underlying CSV file
func (tbl *Table) WriteHeader(hdr string) error {
	if !strings.HasSuffix(hdr, "\n") {
		hdr += "\n"
	}
	_, err := tbl.f.WriteString(hdr)
	return err
}

// WriteRow writes the data into the columns at the current row.
func (tbl *Table) WriteRow(args ...interface{}) error {
	var err error
	if tbl.Writer == nil {
		return fmt.Errorf("csvutil: Table is not in write mode")
	}

	switch len(args) {
	case 0:
		return fmt.Errorf("csvutil: Table.WriteRow needs at least one argument")

	case 1:
		// maybe special case: struct?
		rv := reflect.Indirect(reflect.ValueOf(args[0]))
		rt := rv.Type()
		switch rt.Kind() {
		case reflect.Struct:
			err = tbl.writeStruct(rv)
			return err
		}
	}

	err = tbl.write(args...)
	if err != nil {
		return err
	}

	return err
}

func (tbl *Table) write(args ...interface{}) error {
	rec := make([]string, len(args))
	for i, arg := range args {
		rv := reflect.Indirect(reflect.ValueOf(arg))
		rt := rv.Type()
		switch rt.Kind() {
		case reflect.Bool:
			rec[i] = strconv.FormatBool(rv.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rec[i] = strconv.FormatInt(rv.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rec[i] = strconv.FormatUint(rv.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			rec[i] = strconv.FormatFloat(rv.Float(), 'g', -1, rt.Bits())
		case reflect.String:
			rec[i] = rv.String()
		default:
			return fmt.Errorf("csvutil: invalid type (%[1]T) %[1]v (kind=%[2]v)", arg, rt.Kind())
		}
	}
	return tbl.Writer.Write(rec)
}

func (tbl *Table) writeStruct(rv reflect.Value) error {
	rt := rv.Type()
	args := make([]interface{}, rt.NumField())
	for i := range args {
		args[i] = rv.Field(i).Interface()
	}

	return tbl.write(args...)
}

// Rows is an iterator over an interval of rows inside a CSV file.
type Rows struct {
	tbl    *Table
	i      int64    // number of rows iterated over
	n      int64    // number of rows this iterator iters over
	inc    int64    // number of rows to increment by at each iteration
	cur    int64    // current row index
	record []string // last read record
	closed bool
	err    error // last error
}

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit Close.
func (rows *Rows) Err() error {
	return rows.err
}

// Close closes the Rows, preventing further enumeration.
// Close is idempotent and does not affect the result of Err.
func (rows *Rows) Close() error {
	if rows.closed {
		return nil
	}
	rows.closed = true
	rows.tbl = nil
	return nil
}

// NumFields returns the number of fields in the current CSV-record.
// NumFields assumes Rows.Next() has been called at least once.
func (rows *Rows) NumFields() int {
	return len(rows.record)
}

// Fields returns the raw string values of the fields of the current CSV-record.
// Fields assumes Rows.Next() has been called at least once.
func (rows *Rows) Fields() []string {
	fields := make([]string, len(rows.record))
	copy(fields, rows.record)
	return fields
}

// Scan copies the columns in the current row into the values pointed at by
// dest.
// dest can be either:
// - a pointer to a struct value (whose fields will be filled with column values)
// - a slice of values
func (rows *Rows) Scan(dest ...interface{}) error {
	var err error
	defer func() {
		rows.err = err
	}()

	switch len(dest) {
	case 0:
		err = fmt.Errorf("csvutil: Rows.Scan needs at least one argument")
		return err

	case 1:
		// maybe special case: struct?
		rv := reflect.ValueOf(dest[0]).Elem()
		rt := rv.Type()
		switch rt.Kind() {
		case reflect.Struct:
			err = rows.scanStruct(rv)
			return err
		}
	}

	err = rows.scan(dest...)
	return err
}

func (rows *Rows) scan(args ...interface{}) error {
	var err error
	n := min(len(rows.record), len(args))
	for i := 0; i < n; i++ {
		rec := rows.record[i]
		rv := reflect.ValueOf(args[i]).Elem()
		rt := reflect.TypeOf(args[i]).Elem()
		switch rt.Kind() {
		case reflect.Bool:
			v, err := strconv.ParseBool(rec)
			if err != nil {
				return err
			}
			rv.SetBool(v)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v, err := strconv.ParseInt(rec, 10, rt.Bits())
			if err != nil {
				return err
			}
			rv.SetInt(v)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v, err := strconv.ParseUint(rec, 10, rt.Bits())
			if err != nil {
				return err
			}
			rv.SetUint(v)

		case reflect.Float32, reflect.Float64:
			v, err := strconv.ParseFloat(rec, rt.Bits())
			if err != nil {
				return err
			}
			rv.SetFloat(v)

		case reflect.String:
			rv.SetString(rec)

		default:
			return fmt.Errorf("csvutil: invalid type (%T) %q (kind=%v)", rv.Interface(), rec, rt.Kind())
		}
	}

	return err
}

func (rows *Rows) scanStruct(rv reflect.Value) error {
	rt := rv.Type()
	args := make([]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		args[i] = rv.Field(i).Addr().Interface()
	}
	return rows.scan(args...)
}

func (rows *Rows) skip(n int64) error {
	var err error
	for i := int64(0); i < n; i++ {
		_, err = rows.tbl.Reader.Read()
		if err != nil {
			return err
		}
		rows.cur++
	}
	return err
}

// Next prepares the next result row for reading with the Scan method.
// It returns true on success, false if there is no next result row.
// Every call to Scan, even the first one, must be preceded by a call to Next.
func (rows *Rows) Next() bool {
	if rows.closed {
		return false
	}
	if rows.err != nil {
		return false
	}
	next := rows.i < rows.n
	rows.cur += rows.inc
	rows.i += rows.inc
	if !next {
		rows.err = rows.Close()
		return next
	}

	var err error
	rows.record, err = rows.tbl.Reader.Read()
	if err != nil {
		rows.err = err
		return false
	}

	return next
}
