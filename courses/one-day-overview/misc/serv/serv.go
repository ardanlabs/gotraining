package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":8888"

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello", req.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening on", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
