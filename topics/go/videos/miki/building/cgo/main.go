package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheck() error {
	// TODO
	return nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := healthCheck(); err != nil {
		log.Printf("error: %s", err)
		http.Error(w, "health check failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
