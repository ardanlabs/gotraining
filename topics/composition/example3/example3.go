// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/avOqgN0O4v

// Sample program demonstrating interface composition.
package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// =============================================================================

// Data is the structure of the data we are copying.
type Data struct {
	Line string
}

// =============================================================================

// Puller declares behavior for pulling data.
type Puller interface {
	Pull(d *Data) error
}

// Storer declares behavior for storing data.
type Storer interface {
	Store(d Data)
}

// PullStorer declares behavior for both pulling and storing.
type PullStorer interface {
	Puller
	Storer
}

// =============================================================================

// Xenia is a system we need to pull data from.
type Xenia struct{}

// Pull knows how to pull data out of Xenia.
func (Xenia) Pull(d *Data) error {
	if rand.Intn(10) == 5 {
		return io.EOF
	}

	d.Line = "Data"
	fmt.Println("In:", d.Line)

	return nil
}

// Pillar is a system we need to store data into.
type Pillar struct{}

// Store knows how to store data into Pillar.
func (Pillar) Store(d Data) {
	fmt.Println("Out:", d.Line)
}

// =============================================================================

// System wraps Xenia and Pillar together into a single system.
type System struct {
	Xenia
	Pillar
}

// =============================================================================

// IO provides support to copy bulk data.
type IO struct{}

// pull knows how to pull bulks of data from any Puller.
func (IO) pull(p Puller, data []Data) error {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return err
		}
	}

	return nil
}

// store knows how to store bulks of data from any Storer.
func (IO) store(s Storer, data []Data) {
	for _, d := range data {
		s.Store(d)
	}
}

// Copy knows how to pull and store data from any System.
func (io IO) Copy(ps PullStorer, batch int) {
	for {
		data := make([]Data, batch)
		if err := io.pull(ps, data); err != nil {
			return
		}

		io.store(ps, data)
	}
}

// =============================================================================

// main is the entry point for the application.
func main() {

	// Initialize the system for use.
	sys := System{
		Xenia:  Xenia{},
		Pillar: Pillar{},
	}

	var io IO
	io.Copy(&sys, 3)
}
