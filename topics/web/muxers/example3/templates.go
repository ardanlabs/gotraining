package main

import (
	"html/template"
	"io"
	"path"
	"path/filepath"
	"runtime"

	"github.com/labstack/echo"
)

type renderer struct {
	*template.Template
}

func (r renderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return r.ExecuteTemplate(w, name, data)
}

var templates = renderer{}

func init() {
	pattern := filepath.Join(currentDir(), "templates", "*.html")
	templates.Template = template.Must(template.ParseGlob(pattern))
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
