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

func Exec(w io.Writer) error {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		return err
	}
	return t.Execute(w, nil)
}

func main() {
	log.Fatal(Exec(os.Stdout))
}
