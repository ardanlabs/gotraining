// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csvutil

import (
	"fmt"
	"log"
)

func Example_readStruct() {
	fname := "testdata/simple.csv"
	tbl, err := Open(fname)
	if err != nil {
		log.Fatalf("could not open %s: %v\n", fname, err)
	}
	defer tbl.Close()
	tbl.Reader.Comma = ';'
	tbl.Reader.Comment = '#'

	rows, err := tbl.ReadRows(0, 10)
	if err != nil {
		log.Fatalf("could read rows [0, 10): %v\n", err)
	}
	defer rows.Close()

	irow := 0
	for rows.Next() {
		data := struct {
			I int
			F float64
			S string
		}{}
		err = rows.Scan(&data)
		if err != nil {
			log.Fatalf("error reading row %d: %v\n", irow, err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func Example_readSlice() {
	fname := "testdata/simple.csv"
	tbl, err := Open(fname)
	if err != nil {
		log.Fatalf("could not open %s: %v\n", fname, err)
	}
	defer tbl.Close()
	tbl.Reader.Comma = ';'
	tbl.Reader.Comment = '#'

	rows, err := tbl.ReadRows(0, 10)
	if err != nil {
		log.Fatalf("could read rows [0, 10): %v\n", err)
	}
	defer rows.Close()

	irow := 0
	for rows.Next() {
		var (
			I int
			F float64
			S string
		)
		err = rows.Scan(&I, &F, &S)
		if err != nil {
			log.Fatalf("error reading row %d: %v\n", irow, err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func Example_writeStruct() {
	fname := "out.csv"
	tbl, err := Create(fname)
	if err != nil {
		log.Fatalf("could not create %s: %v\n", fname, err)
	}
	defer tbl.Close()
	tbl.Writer.Comma = ';'

	err = tbl.WriteHeader("## a simple set of data: int64;float64;string\n")
	if err != nil {
		log.Fatalf("error writing header: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		data := struct {
			I int
			F float64
			S string
		}{
			I: i,
			F: float64(i),
			S: fmt.Sprintf("str-%d", i),
		}
		err = tbl.WriteRow(data)
		if err != nil {
			log.Fatalf("error writing row %d: %v\n", i, err)
		}
	}

	err = tbl.Close()
	if err != nil {
		log.Fatalf("error closing table: %v\n", err)
	}
}

func Example_writeSlice() {
	fname := "out.csv"
	tbl, err := Create(fname)
	if err != nil {
		log.Fatalf("could not create %s: %v\n", fname, err)
	}
	defer tbl.Close()
	tbl.Writer.Comma = ';'

	err = tbl.WriteHeader("## a simple set of data: int64;float64;string\n")
	if err != nil {
		log.Fatalf("error writing header: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		var (
			f = float64(i)
			s = fmt.Sprintf("str-%d", i)
		)
		err = tbl.WriteRow(i, f, s)
		if err != nil {
			log.Fatalf("error writing row %d: %v\n", i, err)
		}
	}

	err = tbl.Close()
	if err != nil {
		log.Fatalf("error closing table: %v\n", err)
	}
}
