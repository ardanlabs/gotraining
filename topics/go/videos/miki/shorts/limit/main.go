package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	maxSize = 1 << 20 // 1MB
)

func handler(w http.ResponseWriter, r *http.Request) {
	rdr := http.MaxBytesReader(w, r.Body, maxSize)
	defer rdr.Close()
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "read %d bytes\n", len(data))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("INFO: max request size = %d", maxSize)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
