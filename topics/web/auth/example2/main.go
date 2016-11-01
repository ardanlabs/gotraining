package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

var templates = template.New("HTML")

func init() {
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))
	var err error
	templates, err = templates.New("index").Parse(indexHTML)
	if err != nil {
		log.Fatal(err)
	}
	templates, err = templates.New("user").Parse(userHTML)
	if err != nil {
		log.Fatal(err)
	}
}

func CallbackHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	templates.ExecuteTemplate(res, "user", user)
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "index", nil)
}

func App() http.Handler {
	goth.UseProviders(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://127.0.0.1:3000/auth/github/callback"))

	p := pat.New()
	p.Get("/auth/{provider}/callback", CallbackHandler)
	p.Get("/auth/{provider}", gothic.BeginAuthHandler)
	p.Get("/", IndexHandler)

	return p
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}

var indexHTML = `<p><a href="/auth/github">Log in with Github</a></p>`

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
<p>RefreshToken: {{.RefreshToken}}</p>
`
