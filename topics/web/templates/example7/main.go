package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func App() http.Handler {
	m := http.NewServeMux()

	box := rice.MustFindBox("./static")
	assets := http.StripPrefix("/assets/", http.FileServer(box.HTTPBox()))
	m.Handle("/assets/", assets)

	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		b, _ := box.Bytes("index.html")
		res.Write(b)
	})
	return m
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
