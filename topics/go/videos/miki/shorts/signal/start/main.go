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

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	slog.Info("server starting", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server", "error", err)
		os.Exit(1)
	}
}
