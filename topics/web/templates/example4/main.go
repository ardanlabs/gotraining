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
<a href="/foo?email={{.Email}}">{{.Email}}</a>
<script>
	window.user = {{.}};
</script>
`

// User represents user data.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Exec loads and executes the template. The resulting output
// is sent through the Writer.
func Exec(w io.Writer) error {

	// Create a new template giving it a name, then parse
	// the html string into a template.
	t, err := template.New("foo").Parse(html)
	if err != nil {
		return err
	}

	// Create a value of type user.
	u := User{
		Name:  "Mark",
		Email: "mark@example.com",
	}

	// Execute the parsed template writing the output to
	// the writer. Pass the user value for proecssing.
	return t.Execute(w, u)
}

func main() {

	// Execute the template sending the resulting
	// to stdout.
	if err := Exec(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
