# Your first Web Server

In this quick start guide, we will cover the following basics:

- Spin up a webserver
- Basic Routing
- Your first template

## Running the webserver

The first thing we are going to do is run a webserver. We can do this in just a few lines of code.

Open a file  called `webserver.go` and add the following content:

```go
// https://play.golang.org/p/x4iJhctz8e

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
```

To run this issue the following command:

```sh
go run webserver.go
```

Then open web browser on [http://localhost:8080](http://localhost:8080)

As you can see, we are just writing out a basic "Hello world" followed by the path.

Now go to this url: [http://localhost:8080/foo](http://localhost:8080/foo)

As you can see, this webserver will respond to any route you give it.  We really don't want that so let's change the code as follows:

```go
// https://play.golang.org/p/CUnPy2CKqI

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	fmt.Fprintf(w, "Hello World from %q", html.EscapeString(r.URL.Path))
}
```

You will notice that we are using our first `external` package called `github.com/gorilla/mux`.

To make sure this is on your computer and allow the program to compile you need to `get` it with the following command:

```sh
go get github.com/gorilla/mux
```

Now, we can fire up the server again:

```sh
go run webserver.go
```

Now open web browser on [http://localhost:8080](http://localhost:8080).  As you can see nothing new is really going on.


However, if you go to [http://localhost:8080/foo](http://localhost:8080/foo), you now get a page not found, as we expected


## Templates

Now we need to make a basic template.  To do so, we will use the [http://golang.org/pkg/html/template/](html/template) package.

Change the program as follows:

```go
// https://play.golang.org/p/N5c1LMZWe_

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
```

### Summary

Congratulations, you have your first basic webserver!


