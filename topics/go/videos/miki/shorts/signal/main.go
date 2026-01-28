package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
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

	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("server run", "error", err)
		os.Exit(1)
	}
}
