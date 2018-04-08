// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package csvdriver registers a database/sql/driver.Driver implementation for CSV files.
package csvdriver // import "go-hep.org/x/hep/csvutil/csvdriver"

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	_ "github.com/cznic/ql/driver"
)

var (
	_ driver.Driver  = (*csvDriver)(nil)
	_ driver.Conn    = (*csvConn)(nil)
	_ driver.Execer  = (*csvConn)(nil)
	_ driver.Queryer = (*csvConn)(nil)
	_ driver.Tx      = (*csvConn)(nil)
)

// Conn describes how a connection to the CSV-driver should be established.
type Conn struct {
	File    string      `json:"file"`    // name of the file to be open
	Mode    int         `json:"mode"`    // r/w mode (default: read-only)
	Perm    os.FileMode `json:"perm"`    // file permissions
	Comma   rune        `json:"comma"`   // field delimiter (default: ',')
	Comment rune        `json:"comment"` // comment character for start of line (default: '#')
	Header  bool        `json:"header"`  // whether the CSV-file has a column header
	Names   []string    `json:"names"`   // column names
}

func (c *Conn) setDefaults() {
	if c.Mode == 0 {
		c.Mode = os.O_RDONLY
		c.Perm = 0
	}
	if c.Comma == 0 {
		c.Comma = ','
	}
	if c.Comment == 0 {
		c.Comment = '#'
	}
	return
}

func (c Conn) toJSON() (string, error) {
	c.setDefaults()
	buf, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(buf), err
}

// Open opens a database connection with the CSV driver.
func (c Conn) Open() (*sql.DB, error) {
	c.setDefaults()
	str, err := c.toJSON()
	if err != nil {
		return nil, err
	}
	return sql.Open("csv", str)
}

// Open is a CSV-driver helper function for sql.Open.
//
// It opens a database connection to csvdriver.
func Open(name string) (*sql.DB, error) {
	c := Conn{File: name, Mode: os.O_RDONLY, Perm: 0}
	return c.Open()
}

// Create is a CSV-driver helper function for sql.Open.
//
// It creates a new CSV file, connected via the csvdriver.
func Create(name string) (*sql.DB, error) {
	c := Conn{
		File: name,
		Mode: os.O_RDWR | os.O_CREATE | os.O_TRUNC,
		Perm: 0666,
	}
	return c.Open()
}

// csvDriver implements the interface required by database/sql/driver.
type csvDriver struct {
	dbs map[string]*csvConn
	mu  sync.Mutex
}

// Open returns a new connection to the database.
// The name is a string in a driver-specific format.
//
// Open may return a cached connection (one previously
// closed), but doing so is unnecessary; the sql package
// maintains a pool of idle connections for efficient re-use.
//
// The returned connection is only used by one goroutine at a
// time.
func (drv *csvDriver) Open(cfg string) (driver.Conn, error) {
	c := Conn{}
	if strings.HasPrefix(cfg, "{") {
		err := json.Unmarshal([]byte(cfg), &c)
		if err != nil {
			return nil, err
		}
	} else {
		c.File = cfg
		c.setDefaults()
	}

	doImport := false
	_, err := os.Lstat(c.File)
	if err == nil {
		doImport = true
	}

	drv.mu.Lock()
	defer drv.mu.Unlock()
	if drv.dbs == nil {
		drv.dbs = make(map[string]*csvConn)
	}
	conn := drv.dbs[c.File]
	if conn == nil {
		var f *os.File
		switch {
		case strings.HasPrefix(c.File, "http://"), strings.HasPrefix(c.File, "https://"):
			// FIXME(sbinet: check that c.Mode makes sense (ie: only reading)
			resp, err := http.Get(c.File)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			// FIXME(sbinet): devise a mechanism to remove that temporary file
			// when we close the connection.
			f, err = ioutil.TempFile("", "csvdriver-")
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(f, resp.Body)
			if err != nil {
				return nil, err
			}
			_, err = f.Seek(0, io.SeekStart)
			if err != nil {
				return nil, err
			}
			doImport = true
		default:
			// local file
			f, err = os.OpenFile(c.File, c.Mode, c.Perm)
			if err != nil {
				return nil, err
			}
		}
		conn = &csvConn{
			f:    f,
			cfg:  c,
			drv:  drv,
			refs: 0,
		}

		err = conn.initDB()
		if err != nil {
			return nil, err
		}

		if doImport {
			err = conn.importCSV()
			if err != nil {
				return nil, err
			}
		}
		drv.dbs[c.File] = conn
	}
	conn.refs++

	return conn, err
}

type csvConn struct {
	f    *os.File
	cfg  Conn
	drv  *csvDriver
	refs int

	conn  driver.Conn
	exec  driver.Execer
	query driver.Queryer
	tx    driver.Tx
}

func (conn *csvConn) initDB() error {
	c, err := qlopen(conn.cfg.File)
	if err != nil {
		return err
	}

	conn.conn = c
	conn.exec = c.(driver.Execer)
	conn.query = c.(driver.Queryer)
	return nil
}

// Prepare returns a prepared statement, bound to this connection.
func (conn *csvConn) Prepare(query string) (driver.Stmt, error) {
	return conn.conn.Prepare(query)
}

// Close invalidates and potentially stops any current
// prepared statements and transactions, marking this
// connection as no longer in use.
//
// Because the sql package maintains a free pool of
// connections and only calls Close when there's a surplus of
// idle connections, it shouldn't be necessary for drivers to
// do their own connection caching.
func (conn *csvConn) Close() error {
	if conn.refs > 1 {
		conn.refs--
		return nil
	}
	var err error
	defer conn.f.Close()

	// FIXME(sbinet) write-back to file if needed.
	// err = conn.exportCSV()

	err = conn.conn.Close()
	if err != nil {
		return err
	}

	err = conn.f.Close()
	if err != nil {
		return err
	}

	conn.drv.mu.Lock()
	if conn.refs == 1 {
		delete(conn.drv.dbs, conn.f.Name())
	}
	conn.refs = 0
	conn.drv.mu.Unlock()

	return err
}

// Begin starts and returns a new transaction.
func (conn *csvConn) Begin() (driver.Tx, error) {
	tx, err := conn.conn.Begin()
	if err != nil {
		return nil, err
	}
	conn.tx = tx
	return tx, err
}

func (conn *csvConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	return conn.exec.Exec(query, args)
}

func (conn *csvConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	rows, err := conn.query.Query(query, args)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (conn *csvConn) Commit() error {
	if conn.tx == nil {
		return fmt.Errorf("csvdriver: commit while not in transaction")
	}
	err := conn.tx.Commit()
	conn.tx = nil
	return err
}

func (conn *csvConn) Rollback() error {
	if conn.tx == nil {
		return fmt.Errorf("csvdriver: rollback while not in transaction")
	}
	err := conn.tx.Rollback()
	conn.tx = nil
	return err
}

func qlopen(name string) (driver.Conn, error) {
	conn, err := qldrv.Open("memory://" + name)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

var (
	qldrv driver.Driver
)

func init() {
	sql.Register("csv", &csvDriver{})

	db, err := sql.Open("ql", "memory:///dev/null")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	qldrv = db.Driver()
}
