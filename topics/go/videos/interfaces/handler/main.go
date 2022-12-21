package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("ERROR: bad method (%s) from %s", r.Method, r.RemoteAddr)
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	// FIXME: Do real health check
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	addr := ":8080"
	log.Printf("INFO: server starting on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}
