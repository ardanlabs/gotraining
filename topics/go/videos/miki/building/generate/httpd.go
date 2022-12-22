package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(QuoteDB))
	json.NewEncoder(w).Encode(QuoteDB[i])
}

func main() {
	http.HandleFunc("/quote", quoteHandler)

	addr := ":8080"
	log.Printf("start starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
