package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/time/rate"
)

// 100/sec, burst of 200
var limiter = rate.NewLimiter(rate.Limit(100), 200)

func rated(w http.ResponseWriter, r *http.Request) {
	if !limiter.Allow() {
		status := http.StatusTooManyRequests
		http.Error(w, "too fast", status)
		return
	}
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", rated)

	addr := os.Getenv("RATE_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("INFO: server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
