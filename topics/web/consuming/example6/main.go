package main

import (
	"log"
	"net/http"

	jose "github.com/dvsekhvalnov/jose2go"
)

var SharedSecret = []byte("some shared secret")

func App() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		s := req.Header.Get("x-signature")
		if s == "" {
			res.WriteHeader(http.StatusPreconditionRequired)
			return
		}

		payload, _, err := jose.Decode(s, SharedSecret)
		if err != nil || payload == "" {
			res.WriteHeader(http.StatusPreconditionFailed)
			return
		}
		res.WriteHeader(200)
		res.Write([]byte(payload))
	})
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
