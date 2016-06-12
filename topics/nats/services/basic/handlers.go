// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
//
// Code provided by Kelsey Hightower: https://github.com/kelseyhightower/app
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// HelloHandler is the core API this service provides.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hello",
	}

	json.NewEncoder(w).Encode(response)
	return
}

//==============================================================================

// JWTAuthHandler handles authentication for the service.
func JWTAuthHandler(h http.HandlerFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		prf := func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		}

		token, err := jwt.ParseFromRequest(r, prf)
		if err != nil || !token.Valid {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}

	return f
}

//==============================================================================

// versionHandler is providing version support.
type versionHandler struct {
	version string
}

// VersionHandler returns a Handler interface value that contains
// the versionHandler concrete type support needed for versioning.
func VersionHandler(version string) http.Handler {
	return &versionHandler{
		version: version,
	}
}

// ServeHTTP is the handler called when version information is request.
func (h *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Version string `json:"version"`
	}{
		Version: h.version,
	}

	json.NewEncoder(w).Encode(response)
	return
}

//==============================================================================

// LoggingHandler is middleware that provides request logging support.
func LoggingHandler(h http.Handler) http.Handler {
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := "%s - - [%s] \"%s %s %s\" %s\n"
		log.Printf(format, r.RemoteAddr, time.Now().Format(time.RFC1123),
			r.Method, r.URL.Path, r.Proto, r.UserAgent())

		h.ServeHTTP(w, r)
	})

	return f
}
