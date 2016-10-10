package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

var html = `
Hello, World!
`

func Exec(w io.Writer) {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func main() {
	Exec(os.Stdout)
}
