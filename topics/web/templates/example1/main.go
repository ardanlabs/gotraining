// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use create, parse and execute
// a simple template.
package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

// html represents the template.
var html = `
Hello, World!
`

// Exec loads and executes the template. The resulting output
// is sent through the Writer.
func Exec(w io.Writer) error {

	// Create a new template giving it a name, then parse
	// the html string into a template.
	t, err := template.New("foo").Parse(html)
	if err != nil {
		return err
	}

	// Execute the parsed template writing the output to
	// the writer.
	return t.Execute(w, nil)
}

func main() {

	// Execute the template sending the resulting
	// to stdout.
	if err := Exec(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
