// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use create, parse and execute a template
// with simple data processing. This example uses a struct type value
// with a slice and method for generating HTML markup.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// html represents the template asking for data substitution.
var html = `
<h1>{{.FullName}}</h1>
<h2>{{.FullName | upper}}</h2>

Aliases:
<ul>
	{{range $alias := .Aliases -}}
	<li>{{$alias}}</li>
	{{end}}
</ul>
`

// User represents user data.
type User struct {
	First     string
	Last      string
	CreatedAt time.Time
	Aliases   []string
}

// FullName provides template support for formatting
// data within the processingn of the template.
func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.First, u.Last)
}

// Exec loads and executes the template. The resulting output
// is sent through the Writer.
func Exec(w io.Writer) error {

	// A FuncMap is a map of functions that are bound
	// to a name for template processing.
	funcs := template.FuncMap{
		"upper": strings.ToUpper,
		"fdate": func(t time.Time) string {
			return t.Format(time.RubyDate)
		},
	}

	// Create a new template giving it a name, a set of funcs
	// and then parse the html string into a template.
	t, err := template.New("foo").Funcs(funcs).Parse(html)
	if err != nil {
		return err
	}

	// Create a value of type user.
	u := User{
		First:     "Mary",
		Last:      "Smith",
		CreatedAt: time.Now(),
		Aliases:   []string{"Scarface", "MC Skat Kat"},
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
