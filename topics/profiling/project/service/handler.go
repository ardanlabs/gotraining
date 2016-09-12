package service

import (
	"expvar"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/topics/profiling/project/search"
	"github.com/pborman/uuid"
)

// req keeps track of the number of requests.
var req = expvar.NewInt("requests")

// handler handles the search route processing.
func handler(w http.ResponseWriter, r *http.Request) {
	uid := uuid.New()

	log.Printf("%s : handler : Started : Method[%s] URL[%s]\n", uid, r.Method, r.URL)

	// Add a new counter for monitoring.
	req.Add(1)

	// Capture all the form values.
	fv, options := formValues(r)
	log.Printf("%s : handler : Info : options[%#v]\n", uid, options)

	// If this is a post, perform a search.
	var results []search.Result
	if r.Method == "POST" && options.Term != "" {
		results = search.Submit(uid, options)
	}

	// Render the search page.
	markup := render(fv, results)

	// Write the final markup as the response.
	fmt.Fprint(w, string(markup))

	log.Printf("%s : handler : Completed\n", uid)
}

// formValues extracts the form data.
func formValues(r *http.Request) (map[string]interface{}, search.Options) {
	fv := make(map[string]interface{})
	var options search.Options

	fv["term"] = r.FormValue("term")
	options.Term = r.FormValue("term")

	if r.FormValue("cnn") == "on" {
		fv["cnn"] = "checked"
		options.CNN = true
	} else {
		fv["cnn"] = ""
	}

	if r.FormValue("nyt") == "on" {
		fv["nyt"] = "checked"
		options.NYT = true
	} else {
		fv["nyt"] = ""
	}

	if r.FormValue("bbc") == "on" {
		fv["bbc"] = "checked"
		options.BBC = true
	} else {
		fv["bbc"] = ""
	}

	if r.FormValue("first") == "on" {
		fv["first"] = "checked"
		options.First = true
	} else {
		fv["first"] = ""
	}

	return fv, options
}

// render generates the HTML response for this route.
func render(fv map[string]interface{}, results []search.Result) []byte {

	// Generate the markup for the results template.
	if results != nil {
		vars := map[string]interface{}{"Items": results}
		markup := executeTemplate("results", vars)
		fv["Results"] = template.HTML(string(markup))
	}

	// Generate the markup for the search template.
	markup := executeTemplate("search", fv)

	// Generate the final markup with the layout template.
	vars := map[string]interface{}{"LayoutContent": template.HTML(string(markup))}
	return executeTemplate("layout", vars)
}
