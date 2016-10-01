package main

import (
	"log"
	"net/http"
)

var apis map[string]http.Handler

const defaultAPIVersion = "2"

func init() {
	apis = map[string]http.Handler{
		"1": V1(),
		"2": V2(),
	}
}

func V1() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("v1"))
	})
	return r
}

func V2() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("v2"))
	})
	return r
}

func App() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/api/", func(res http.ResponseWriter, req *http.Request) {
		v := req.Header.Get("x-version")
		h := apis[v]
		if h == nil {
			h = apis[defaultAPIVersion]
		}
		http.StripPrefix("/api", h).ServeHTTP(res, req)
	})
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
