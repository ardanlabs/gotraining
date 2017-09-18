// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csvdriver_test

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"go-hep.org/x/hep/csvutil/csvdriver"
)

func testDB(t *testing.T, conn csvdriver.Conn, vars string) {
	db, err := conn.Open()
	if err != nil {
		t.Errorf("%s: error opening CSV file", conn.File)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Errorf("%s: error starting tx: %v", conn.File, err)
		return
	}
	defer tx.Commit()

	done := make(chan error)
	go func() {
		done <- db.Ping()
	}()

	select {
	case <-time.After(2 * time.Second):
		t.Errorf("%s: ping timeout", conn.File)
		return
	case err := <-done:
		if err != nil {
			t.Errorf("%s: error pinging db: %v\n", conn.File, err)
			return
		}
	}

	rows, err := tx.Query("select " + vars + " from csv order by id();")
	if err != nil {
		t.Errorf("%s: error querying db: %v\n", conn.File, err)
		return
	}
	defer rows.Close()

	type dataType struct {
		i int64
		f float64
		s string
	}

	var got []dataType
	for rows.Next() {
		var data dataType
		err = rows.Scan(&data.i, &data.f, &data.s)
		if err != nil {
			t.Errorf("%s: error scanning db: %v\n", conn.File, err)
			return
		}
		got = append(got, data)
	}

	err = rows.Close()
	if err != nil {
		t.Errorf("%s: error closing rows: %v\n", conn.File, err)
		return
	}

	err = db.Close()
	if err != nil {
		t.Errorf("%s: error closing db: %v\n", conn.File, err)
		return
	}

	want := []dataType{
		{0, 0, "str-0"},
		{1, 1, "str-1"},
		{2, 2, "str-2"},
		{3, 3, "str-3"},
		{4, 4, "str-4"},
		{5, 5, "str-5"},
		{6, 6, "str-6"},
		{7, 7, "str-7"},
		{8, 8, "str-8"},
		{9, 9, "str-9"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got=\n%v\nwant=\n%v\n", conn.File, got, want)
		return
	}
}

func TestOpen(t *testing.T) {
	for _, test := range []struct {
		c csvdriver.Conn
		q string
	}{
		{
			c: csvdriver.Conn{
				File:    "testdata/simple.csv",
				Comment: '#', Comma: ';',
			},
			q: "var1, var2, var3",
		},
		{
			c: csvdriver.Conn{
				File:    "testdata/simple.csv",
				Comment: '#', Comma: ';',
				Names: []string{"v1", "v2", "v3"},
			},
			q: "v1, v2, v3",
		},
		{
			c: csvdriver.Conn{
				File:    "testdata/simple-with-comment.csv",
				Comment: '#', Comma: ';',
			},
			q: "var1, var2, var3",
		},
		{
			c: csvdriver.Conn{
				File:    "testdata/simple-with-comment.csv",
				Comment: '#', Comma: ';',
				Names: []string{"v1", "v2", "v3"},
			},
			q: "v1, v2, v3",
		},
		{
			c: csvdriver.Conn{
				File:    "testdata/simple-with-header.csv",
				Comment: '#', Comma: ';',
				Header: true,
			},
			q: "i64, f64, str",
		},
		{
			c: csvdriver.Conn{
				File: "testdata/simple-with-header.csv", Comment: '#', Comma: ';',
				Header: true,
				Names:  []string{"var1", "var2", "var3"},
			},
			q: "var1, var2, var3",
		},
	} {
		testDB(t, test.c, test.q)
	}
}

func TestQL(t *testing.T) {
	db, err := sql.Open("ql", "memory://out-create-ql.csv")
	if err != nil {
		t.Fatalf("error creating CSV-QL file: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("error pinging db: %v\n", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("error starting transaction: %v\n", err)
	}
	defer tx.Commit()

	_, err = tx.Exec("create table csv (var1 int64, var2 float64, var3 string);")
	if err != nil {
		t.Fatalf("error creating table: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		f := float64(i)
		s := fmt.Sprintf("str-%d", i)
		_, err = tx.Exec("insert into csv values($1,$2,$3);", i, f, s)
		if err != nil {
			t.Fatalf("error inserting row %d: %v\n", i+1, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		t.Fatalf("error committing transaction: %v\n", err)
	}
}

func TestCreate(t *testing.T) {
	const fname = "testdata/out-create.csv"
	defer os.Remove(fname)

	db, err := csvdriver.Create(fname)
	if err != nil {
		t.Fatalf("error creating CSV file: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("error pinging db: %v\n", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("error starting transaction: %v\n", err)
	}
	defer tx.Commit()

	_, err = tx.Exec("create table csv (var1 int64, var2 float64, var3 string);")
	if err != nil {
		t.Fatalf("error creating table: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		f := float64(i)
		s := fmt.Sprintf("str-%d", i)
		_, err = tx.Exec("insert into csv values($1,$2,$3);", i, f, s)
		if err != nil {
			t.Fatalf("error inserting row %d: %v\n", i+1, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		t.Fatalf("error committing transaction: %v\n", err)
	}
}
