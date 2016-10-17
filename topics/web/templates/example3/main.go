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

func Exec(w io.Writer) error {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		return err
	}
	return t.Execute(w, User{Name: "Mark"})
}

func main() {
	log.Fatal(Exec(os.Stdout))
}
