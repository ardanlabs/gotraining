package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

var html = `
<a href="/foo?email={{.Email}}">{{.Email}}</a>
<script>
	window.user = {{.}};
</script>
`

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Exec(w io.Writer) {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, User{Name: "Mark", Email: "mark@example.com"})
}

func main() {
	Exec(os.Stdout)
}
