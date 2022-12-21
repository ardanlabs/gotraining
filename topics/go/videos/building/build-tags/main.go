package main

import (
	"fmt"
	"log"
	"net/http"
)

func byHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, Founder)
}

func main() {
	http.HandleFunc("/by", byHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
