// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use create, parse and execute
// a template with simple data processing. This example uses
// a struct type value and generates HTML markup.
package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

// html represents the template asking for data substitution.
var html = `
{{ . }}
{{ safe . }}
`

// Exec loads and executes the template. The resulting output
// is sent through the Writer.
func Exec(w io.Writer) error {

	// Create a new template giving it a name, then parse
	// the html string into a template.
	t := template.New("foo")
	t, err := t.Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).Parse(html)

	if err != nil {
		return err
	}

	// Execute the parsed template writing the output to
	// the writer. Pass the user value for proecssing.
	return t.Execute(w, `<script>alert("boo!")</script>`)
}

func main() {

	// Execute the template sending the resulting
	// to stdout.
	if err := Exec(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
