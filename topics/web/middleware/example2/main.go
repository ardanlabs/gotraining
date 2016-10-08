package main

import (
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

func fooHeader(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	res.Header().Set("foo", "bar")
	next(res, req)
}

func logger(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()
	defer func() {
		d := time.Now().Sub(start)
		log.Printf("%s %s %s", req.Method, req.URL.Path, d)
	}()
	next(res, req)
}

func App() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})

	n := negroni.New()
	n.Use(negroni.HandlerFunc(logger))
	n.Use(negroni.HandlerFunc(fooHeader))
	n.UseHandler(m)
	return n
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
