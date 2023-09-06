package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type CtxKeyType int

const IDCtxKey CtxKeyType = 3

// Middleware that adds request ID to request context
func RequestID(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		ctx := context.WithValue(r.Context(), IDCtxKey, id)
		r = r.Clone(ctx)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	rid := r.Context().Value(IDCtxKey)
	fmt.Fprintf(w, "ID: %v\n", rid)
}

func main() {
	h := RequestID(http.HandlerFunc(handler))
	http.Handle("/", h)

	addr := os.Getenv("CTX_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("INFO: server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
