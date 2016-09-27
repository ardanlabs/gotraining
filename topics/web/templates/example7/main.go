package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	m := http.NewServeMux()

	box := rice.MustFindBox("./static")
	assets := http.StripPrefix("/assets/", http.FileServer(box.HTTPBox()))
	m.Handle("/assets/", assets)

	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		b, _ := box.Bytes("index.html")
		res.Write(b)
	})
	log.Fatal(http.ListenAndServe(":3000", m))
}
