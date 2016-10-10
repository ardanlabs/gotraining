package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

var html = `
Hello, {{.}}!
`

func Exec(w io.Writer) {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "World")
}

func main() {
	Exec(os.Stdout)
}
