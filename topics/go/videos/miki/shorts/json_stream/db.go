// Dummy DB package
package main

import (
	"fmt"
	"time"
)

type DB struct{}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) Query() *Rows {
	return &Rows{max: 5}
}

type Rows struct {
	max int
	n   int
}

func (r *Rows) Next() bool {
	r.n++
	return r.n <= r.max
}

func (r *Rows) Event() Event {
	t := time.Now().UTC().Add(-time.Duration((r.max-r.n-200)*527) * time.Millisecond)
	return Event{
		Time:   t,
		User:   "elliot",
		Action: "READ",
		URI:    fmt.Sprintf("file:///reports/sec/%d.txt", r.n),
	}
}

func (r *Rows) Err() error {
	return nil
}
