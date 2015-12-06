// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/nMJIHaNXxm

// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// reader is an interface that defines the act of reading data.
type reader interface {
	read() ([]byte, error)
}

// file defines a system file.
type file struct {
	name string
}

// read implements the reader interface for a file.
func (f file) read() ([]byte, error) {
	return []byte(`<rss><channel><title>Going Go Programming</title></channel></rss>`), nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (p pipe) read() ([]byte, error) {
	return []byte(`{name: "bill", title: "developer"}`), nil
}

// main is the entry point for the application.
func main() {
	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call the retrieve funcion for each concrete type.
	retrieve(f)
	retrieve(p)
}

// retrieve can read any device and process the data.
func retrieve(r reader) error {
	d, err := r.read()
	if err != nil {
		return err
	}

	fmt.Println(string(d))
	return nil
}
