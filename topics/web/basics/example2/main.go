package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World!")
	})
	log.Panic(http.ListenAndServe(":3000", m))
}
