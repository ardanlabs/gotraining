package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	defer func() {
		slog.Info("Heghlu'meH QaQ jajvam")
	}()

	http.HandleFunc("/", handler)
	const addr = ":8080"
	slog.Info("server starting", "address", addr)
	http.ListenAndServe(addr, nil)
}
