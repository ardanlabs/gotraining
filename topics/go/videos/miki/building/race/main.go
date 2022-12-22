package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	counter int64
)

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintln(w, counter)
}

func main() {
	http.HandleFunc("/", handler)

	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
