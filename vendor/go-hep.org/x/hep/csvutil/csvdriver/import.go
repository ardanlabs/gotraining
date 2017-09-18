// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csvdriver

import (
	"database/sql/driver"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"go-hep.org/x/hep/csvutil"
)

func (conn *csvConn) importCSV() error {
	tbl, err := csvutil.Open(conn.cfg.File)
	if err != nil {
		return err
	}
	defer tbl.Close()
	tbl.Reader.Comma = conn.cfg.Comma
	tbl.Reader.Comment = conn.cfg.Comment

	schema, err := inferSchema(conn, conn.cfg.Header, conn.cfg.Names)
	if err != nil {
		return err
	}

	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	_, err = conn.Exec("create table csv ("+schema.Decl()+")", nil)
	if err != nil {
		return err
	}

	_, err = conn.Exec("create index csv_id on csv (id());", nil)
	if err != nil {
		return err
	}

	beg := int64(0)
	if conn.cfg.Header {
		beg++
	}
	rows, err := tbl.ReadRows(beg, -1)
	if err != nil {
		return err
	}
	defer rows.Close()

	vargs, pargs := schema.Args()
	def := schema.Def()
	insert := "insert into csv values(" + def + ");"
	for rows.Next() {
		err = rows.Scan(pargs...)
		if err != nil {
			return err
		}
		for i, arg := range pargs {
			vargs[i] = reflect.ValueOf(arg).Elem().Interface()
		}
		_, err = conn.Exec(insert, vargs)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func inferSchema(conn *csvConn, header bool, names []string) (schemaType, error) {
	tbl, err := csvutil.Open(conn.cfg.File)
	if err != nil {
		return nil, err
	}
	defer tbl.Close()
	tbl.Reader.Comma = conn.cfg.Comma
	tbl.Reader.Comment = conn.cfg.Comment

	return inferSchemaFromTable(tbl, header, names)
}

func inferSchemaFromTable(tbl *csvutil.Table, header bool, names []string) (schemaType, error) {
	var (
		beg int64 = 0
		end int64 = 1
	)
	if header {
		end++
	}
	rows, err := tbl.ReadRows(beg, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if header {
		if !rows.Next() {
			return nil, rows.Err()
		}
		if len(names) == 0 {
			names = rows.Fields()
		}
	}

	if !rows.Next() {
		return nil, rows.Err()
	}

	return inferSchemaFromFields(rows.Fields(), names)
}

func inferSchemaFromFields(fields []string, names []string) (schemaType, error) {
	if len(names) == 0 {
		names = make([]string, len(fields))
	}
	schema := make(schemaType, len(fields))
	for i, field := range fields {
		var err error
		name := names[i]
		if name == "" {
			name = fmt.Sprintf("var%d", i+1)
		}

		schema[i].n = name
		_, err = strconv.ParseInt(field, 10, 64)
		if err == nil {
			schema[i].v = reflect.ValueOf(int64(0))
			continue
		}

		_, err = strconv.ParseFloat(field, 64)
		if err == nil {
			schema[i].v = reflect.ValueOf(float64(0))
			continue
		}

		schema[i].v = reflect.ValueOf("")
	}
	return schema, nil
}

type schemaType []struct {
	v reflect.Value
	n string
}

func (st *schemaType) Decl() string {
	o := make([]string, 0, len(*st))
	for _, v := range *st {
		n := v.n
		t := v.v.Type().Kind().String()
		o = append(o, n+" "+t)
	}
	return strings.Join(o, ", ")
}

func (st *schemaType) Args() ([]driver.Value, []interface{}) {
	vargs := make([]driver.Value, len(*st))
	pargs := make([]interface{}, len(*st))
	for i, v := range *st {
		ptr := reflect.New(v.v.Type())
		vargs[i] = ptr.Elem().Interface()
		pargs[i] = ptr.Interface()
	}
	return vargs, pargs
}

func (st *schemaType) Def() string {
	o := make([]string, len(*st))
	for i := range *st {
		o[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(o, ", ")
}
