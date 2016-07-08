// Package service : temnplates provides support for using HTML
// based templates for responses.
package service

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// views contains a map of templates for rendering views.
var views = make(map[string]*template.Template)

// init loads the existing templates for use by routing code.
func init() {
	// In order for the endpoint tests to run this needs to be
	// physically located. Trying to avoid configuration for now.
	pwd, _ := os.Getwd()
	loadTemplate("layout", pwd+"/views/basic-layout.html")
	loadTemplate("search", pwd+"/views/search.html")
	loadTemplate("results", pwd+"/views/results.html")
}

// loadTemplate reads the specified template file for use.
func loadTemplate(name string, path string) {
	// Read the html template file.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a template value for this code.
	tmpl, err := template.New(name).Parse(string(data))
	if err != nil {
		log.Fatalln(err)
	}

	// Have we processed this file already?
	if _, exists := views[name]; exists {
		log.Fatalf("Template %s already in use.", name)
	}

	// Store the template for use.
	views[name] = tmpl
}

// executeTemplate executes the specified template with the specified variables.
func executeTemplate(name string, vars map[string]interface{}) []byte {
	markup := new(bytes.Buffer)
	if err := views[name].Execute(markup, vars); err != nil {
		log.Println(err)
		return []byte("Error Processing Template")
	}

	return markup.Bytes()
}
