// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/NXTdWFDJoP

// Sample program to show off Go and check programming environment.
package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

// Response is used to send data back to the client.
type Response struct {
	Greeting string
}

// main is the entry point for the application.
func main() {
	http.HandleFunc("/english", helloEnglish)
	http.HandleFunc("/chinese", helloChinese)

	addr := "localhost:9999"
	log.Printf("starting server on http://%s", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}

// helloEnglish sends a greeting in English.
func helloEnglish(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Response{Greeting: "Hello World"}); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
		return
	}

	simpleLog(w, r)
}

// helloChinese sends a greeting in Chinese.
func helloChinese(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Response{Greeting: "你好世界"}); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
		return
	}

	simpleLog(w, r)
}

// simpleLog is will log a very simple set of information for a request
func simpleLog(w http.ResponseWriter, r *http.Request) {
	// try to parse the remote address
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	uri := r.URL.RequestURI()

	referer := r.Referer()
	if referer == "" {
		referer = "-"
	}

	userAgent := r.UserAgent()
	if userAgent == "" {
		userAgent = "-"
	}

	username := "-"

	// Try to get it from the authorization header if set there.
	if u, _, ok := r.BasicAuth(); ok {
		username = u
	}

	fields := []string{
		host,
		username,
		r.Method,
		uri,
		r.Proto,
		referer,
		userAgent,
	}

	log.Println(strings.Join(fields, " "))
}
