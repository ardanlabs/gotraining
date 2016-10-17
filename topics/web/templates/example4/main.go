package main

import (
	"html/template"
	"io"
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

func Exec(w io.Writer) error {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		return err
	}
	return t.Execute(w, User{Name: "Mark", Email: "mark@example.com"})
}

func main() {
	Exec(os.Stdout)
}
