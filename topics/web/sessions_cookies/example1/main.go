package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const sessionName = "ultimate-web-session"

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func homeHandler(res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, sessionName)
	if err != nil {
		fmt.Fprint(res, err)
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
		fmt.Fprint(res, err)
		return
	}

	req.ParseForm()
	name := req.FormValue("myName")
	session.Values["name"] = name
	session.Save(req, res)
	fmt.Fprintf(res, htmlWithSession, name)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/save", saveHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

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
