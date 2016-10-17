package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const sessionName = "ultimate-web-session"

var store = sessions.NewCookieStore([]byte("something-very-secret"))

var htmlNoSession = `
<html>
  <form action="/save" method="POST">
    <label>What is your name?</label><br>
    <input type="text" name="myName" placeholder="Name goes here">
    <input type="submit" value="Submit">
  </form>
</html>
`

var htmlWithSession = `
<html>
  <h1>Hello %s!</h1>
</html>
`

func App() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", homeHandler)
	m.HandleFunc("/save", saveHandler)
	return m
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, sessionName)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	name := session.Values["name"]
	if name != nil {
		fmt.Fprintf(res, htmlWithSession, name)
	} else {
		fmt.Fprint(res, htmlNoSession)
	}
}

func saveHandler(res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, sessionName)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	req.ParseForm()
	name := req.FormValue("myName")
	session.Values["name"] = name
	session.Save(req, res)
	fmt.Fprintf(res, htmlWithSession, name)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
