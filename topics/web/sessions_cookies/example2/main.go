// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use a cookie in your web app.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// cookieName contains the name of the cookie.
const cookieName = "ultimate-web-cookie"

// htmlWithSession contains the document we will use we
// have a request has not submited state yet.
var htmlNoCookie = `
<html>
    <form action="/save" method="POST">
        <label>What is your name?</label><br>
        <input type="text" name="myName" placeholder="Name goes here">
        <input type="submit" value="Submit">
    </form>
</html>`

// htmlWithSession contains the document we will use we
// have a request has already submited state.
var htmlWithCookie = `
<html>
    <h1>Hello %s!</h1>
</html>`

// App loads the entire API set together for use.
func App() http.Handler {

	// Create a new mux which will process the requests.
	m := http.NewServeMux()

	// Load the two routes.
	m.HandleFunc("/", homeHandler)
	m.HandleFunc("/save", saveHandler)

	return m
}

// homeHandler provides support for the home page route.
func homeHandler(res http.ResponseWriter, req *http.Request) {

	// If there is no cookie in the request provide document
	// to collect data.
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		fmt.Fprint(res, htmlNoCookie)
		return
	}

	// There is a cookie so return the document with
	// the saved state.
	fmt.Fprintf(res, htmlWithCookie, cookie.Value)
}

// saveHandler provides support for save route.
func saveHandler(res http.ResponseWriter, req *http.Request) {

	// Parse the raw query from the URL and update req.Form.
	req.ParseForm()

	// Locate the myName form value.
	name := req.FormValue("myName")

	// Set the expiration date to a year from now.
	expiration := time.Now().Add(365 * 24 * time.Hour)

	// Create a new cookie with the data from the request
	// and the expiration date.
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   name,
		Expires: expiration,
	}

	// Set the cookie into the response.
	http.SetCookie(res, cookie)

	// Return a document with the saved state.
	fmt.Fprintf(res, htmlWithCookie, name)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
