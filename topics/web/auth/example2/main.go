// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to apply basic authentication with the
// goth package for your web request.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

// indexHTML contains the base markup for the index page.
var indexHTML = `
<p><a href="/auth/github">Log in with Github</a></p>`

// userHTML contains the base markup for the user page.
var userHTML = `
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>`

// templates will contain the html template for
// rendering the view.
var templates = template.New("HTML")

func init() {

	// A physical location where sessions will be saved.
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))

	// Load the index page template.
	templates, err := templates.New("index").Parse(indexHTML)
	if err != nil {
		log.Fatal(err)
	}

	// Load the index page template.
	templates, err = templates.New("user").Parse(userHTML)
	if err != nil {
		log.Fatal(err)
	}
}

// callbackHandler handles the rendering of the user page.
func callbackHandler(res http.ResponseWriter, req *http.Request) {

	// Complete the authentication process and fetch all of the
	// basic information about the user from the provider.
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	// Execute the template for this user and respond with
	// the user page.
	templates.ExecuteTemplate(res, "user", user)
}

// indexHandler handles the rendering of the index page.
func indexHandler(res http.ResponseWriter, req *http.Request) {

	// Execute the template and respond with the index page.
	templates.ExecuteTemplate(res, "index", nil)
}

// App loads the API for use.
func App() http.Handler {

	// Create a new Github provider with our connection details.
	goth.UseProviders(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://127.0.0.1:3000/auth/github/callback"))

	// Create a new pat router.
	p := pat.New()

	// Bind the user page handler.
	p.Get("/auth/{provider}/callback", callbackHandler)

	// Bind the authentication route.
	p.Get("/auth/{provider}", gothic.BeginAuthHandler)

	// Bind the index page handler.
	p.Get("/", indexHandler)

	return p
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
