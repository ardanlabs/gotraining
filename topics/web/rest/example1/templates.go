package main

import (
	"html/template"
	"path"
	"path/filepath"
	"runtime"
)

var templates *template.Template

func init() {
	pattern := filepath.Join(currentDir(), "templates", "*.html")
	templates = template.Must(template.ParseGlob(pattern))
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
