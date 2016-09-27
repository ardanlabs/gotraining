package main

import (
	"html/template"
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

func main() {
	t, err := template.New("foo").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, User{Name: "Mark", Email: "mark@example.com"})
}
