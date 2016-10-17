package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const cookieName = "ultimate-web-cookie"

var htmlNoCookie = `
<html>
  <form action="/save" method="POST">
    <label>What is your name?</label><br>
    <input type="text" name="myName" placeholder="Name goes here">
    <input type="submit" value="Submit">
  </form>
</html>
`

var htmlWithCookie = `
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
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		fmt.Fprint(res, htmlNoCookie)
		return
	}

	fmt.Fprintf(res, htmlWithCookie, cookie.Value)
}

func saveHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.FormValue("myName")

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := &http.Cookie{Name: cookieName, Value: name, Expires: expiration}
	http.SetCookie(res, cookie)

	fmt.Fprintf(res, htmlWithCookie, name)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
