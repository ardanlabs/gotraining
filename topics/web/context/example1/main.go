package main

import (
	"context"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

type User struct {
	Username string
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	u := req.Context().Value("current_user").(*User)
	res.Write([]byte(u.Username))
}

func CurrentUserMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	u := &User{"mary-jane"}
	ctx := context.WithValue(req.Context(), "current_user", u)
	req = req.WithContext(ctx)
	next(res, req)
}

func App() http.Handler {
	m := http.NewServeMux()
	n := negroni.New()
	n.UseFunc(CurrentUserMiddleware)
	m.HandleFunc("/", IndexHandler)
	n.UseHandler(m)
	return n
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
