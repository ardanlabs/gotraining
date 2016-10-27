package main

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/urfave/negroni"
)

type User struct {
	Username string
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	u := context.Get(req, "current_user").(*User)
	res.Write([]byte(u.Username))
}

func CurrentUserMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	u := &User{"mary-jane"}
	context.Set(req, "current_user", u)
	next(res, req)
}

func App() http.Handler {
	m := http.NewServeMux()
	n := negroni.New()
	n.UseFunc(CurrentUserMiddleware)
	m.HandleFunc("/", IndexHandler)
	n.UseHandler(context.ClearHandler(m))
	return n
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
