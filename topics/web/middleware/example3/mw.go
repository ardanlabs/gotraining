package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// logger is a middleware that wraps the logger from gorilla/handlers
func logger(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

// fooHeader returns a handler function that will set the `foo` header
// key then call the next handler.
func fooHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("foo", "bar")

		next.ServeHTTP(res, req)
	})
}

// decoderCheck is a middleware that looks for a secret value in a
// request header. If it is not present then it will reject requests.
// A better solution is introduced in the topics/web/auth section.
func decoderCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("X-Decoder-Ring") != "Little Orphan Annie" {
			res.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(res, req)
	})
}
