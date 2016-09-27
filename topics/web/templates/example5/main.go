package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

var html = `
<h1>{{.FullName}}</h1>
<h2>{{.FullName | upper}}</h2>

Aliases:
<ul>
  {{range $alias := .Aliases -}}
    <li>{{$alias}}</li>
  {{end }}
</ul>
`

type User struct {
	First   string
	Last    string
	Aliases []string
}

func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.First, u.Last)
}

func main() {
	funcs := template.FuncMap{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}
	t, err := template.New("foo").Funcs(funcs).Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, User{
		First:   "Mary",
		Last:    "Smith",
		Aliases: []string{"Scarface", "MC Skat Kat"},
	})
}
