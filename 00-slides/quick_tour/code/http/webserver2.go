// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/N5c1LMZWe_

// Program to show how to run a basic web server with routing and templating.
package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// This is a basic struct to hold basic page data variables
type PageData struct {
	Title string
	Body  string
}

func main() {
	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path
	rt.HandleFunc("/", Index)

	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	// Fill out page data for index
	pd := PageData{
		Title: "Index Page",
		Body:  "This is the body of the page.",
	}

	// Render a template with our page data
	tmpl, err := htmlTemplate(pd)

	// If we got an error, write it out and exit
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// All went well, so write out the template
	w.Write([]byte(tmpl))
}

func htmlTemplate(pd PageData) (string, error) {
	// Define a basic HTML template
	html := `<HTML>
	<head><title>{{.Title}}</title></head>
	<body>
	{{.Body}}
	</body>
	</HTML>`

	// Parse the template
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}

	// We need somewhere to write the executed template to
	var out bytes.Buffer

	// Render the template with the data we passed in
	if err := tmpl.Execute(&out, pd); err != nil {
		// If we couldn't render, return a error
		return "", err
	}

	// Return the template
	return out.String(), nil
}
