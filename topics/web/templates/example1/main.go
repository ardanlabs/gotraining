package main

import (
	"html/template"
	"log"
	"os"
)

var html = `
Hello, World!
`

func main() {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, nil)
}
