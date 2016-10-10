package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

var html = `
Hello, {{.Name}}!
`

type User struct {
	Name string
}

func Exec(w io.Writer) {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, User{Name: "Mark"})
}

func main() {
	Exec(os.Stdout)
}
