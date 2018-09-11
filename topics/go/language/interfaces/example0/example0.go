// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that could benefit from polymorphic behavior with interfaces.
package main

import "fmt"

// file defines a system file.
type file struct {
	name string
}

// read implements the reader interface for a file.
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

func main() {

	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call each retrieve function for each concrete type.
	retrieveFile(f)
	retrievePipe(p)
}

// retrieveFile can read from a file and process the data.
func retrieveFile(f file) error {
	data := make([]byte, 100)

	len, err := f.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// retrievePipe can read from a pipe and process the data.
func retrievePipe(p pipe) error {
	data := make([]byte, 100)

	len, err := p.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
