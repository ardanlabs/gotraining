// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package service : index maintains the support for the home page.
package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ArdanStudios/gotraining/web_app/sample/search"
)

// index handles the home page route processing.
func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("service : index : Started : Method[%s]\n", r.Method)

	// Capture all the form values.
	fv, options := formValues(r)
	log.Printf("service : index : Info : options[%#v]\n", options)

	// If this is a post, perform a search.
	var results []search.Result
	if r.Method == "POST" && options.SearchTerm != "" {
		results = search.Submit(options)
	}

	// Render the index page.
	markup := renderIndex(fv, results)

	// Write the final markup as the response.
	fmt.Fprint(w, string(markup))

	log.Println("service : index : Completed")
}

// formValues extracts the form data.
func formValues(r *http.Request) (map[string]interface{}, *search.Options) {
	fv := make(map[string]interface{})
	var options search.Options

	fv["searchterm"] = r.FormValue("searchterm")
	options.SearchTerm = r.FormValue("searchterm")

	if r.FormValue("google") == "on" {
		fv["google"] = "checked"
		options.Google = true
	} else {
		fv["google"] = ""
	}

	if r.FormValue("bing") == "on" {
		fv["bing"] = "checked"
		options.Bing = true
	} else {
		fv["bing"] = ""
	}

	if r.FormValue("blekko") == "on" {
		fv["blekko"] = "checked"
		options.Blekko = true
	} else {
		fv["blekko"] = ""
	}

	if r.FormValue("first") == "on" {
		fv["first"] = "checked"
		options.First = true
	} else {
		fv["first"] = ""
	}

	return fv, &options
}

// renderIndex generates the HTML response for this route.
func renderIndex(fv map[string]interface{}, results []search.Result) []byte {
	// Generate the markup for the results template.
	if results != nil {
		vars := map[string]interface{}{"Items": results}
		markup := executeTemplate("results", vars)
		fv["Results"] = template.HTML(string(markup))
	}

	// Generate the markup for the index template.
	markup := executeTemplate("index", fv)

	// Generate the final markup with the layout template.
	vars := map[string]interface{}{"LayoutContent": template.HTML(string(markup))}
	return executeTemplate("layout", vars)
}
