package main

import (
	"log"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	u, p, ok := req.BasicAuth()
	if !ok {
		http.Error(res, "Not authorized", http.StatusUnauthorized)
		return
	}
	if u != "username" && p != "password" {
		http.Error(res, "Not authorized", http.StatusUnauthorized)
		return
	}
	res.Write([]byte("Welcome Authorized User!"))
}

func App() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", IndexHandler)
	return m
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
