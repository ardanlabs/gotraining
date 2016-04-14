// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/5M0cu4ui2k

// Sample program demonstrating struct composition.
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
func (Pillar) Store(d Data) error {
	fmt.Println("Out:", d.Line)
	return nil
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

// pull knows how to pull bulks of data from Xenia.
func (IO) pull(x *Xenia, data []Data) error {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return err
		}
	}

	return nil
}

// store knows how to store bulks of data from Pillar.
func (IO) store(p *Pillar, data []Data) {
	for _, d := range data {
		p.Store(d)
	}
}

// Copy knows how to pull and store data from the System.
func (io IO) Copy(sys *System, batch int) {
	for {
		data := make([]Data, batch)
		if err := io.pull(&sys.Xenia, data); err != nil {
			return
		}

		io.store(&sys.Pillar, data)
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
