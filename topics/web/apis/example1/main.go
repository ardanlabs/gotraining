package main

import (
	"log"
	"net/http"
)

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
	r.Handle("/api/v1/", http.StripPrefix("/api/v1", V1()))
	r.Handle("/api/v2/", http.StripPrefix("/api/v2", V2()))
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
