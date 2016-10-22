package main

import (
	"log"
	"net/http"
	"time"
)

func fooHeader(hf http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("foo", "bar")
		hf(res, req)
	}
}

func logger(hf http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer func() {
			d := time.Now().Sub(start)
			log.Printf("%s %s %s", req.Method, req.URL.Path, d)
		}()
		hf(res, req)
	}
}

func App() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", logger(fooHeader(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})))
	return m
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
