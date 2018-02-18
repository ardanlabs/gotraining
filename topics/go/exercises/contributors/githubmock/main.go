package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
)

var pathRE = regexp.MustCompile("^/repos/[^/]+/[^/]+/contributors$")
var authRE = regexp.MustCompile("^token [A-Za-z0-9]+$")

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.Method, r.URL.Path)

	// They must provide an auth token.
	if !authRE.MatchString(r.Header.Get("Authorization")) {
		http.Error(w, "Authorization header must be in the form \"token {githubToken}\"", http.StatusUnauthorized)
		return
	}

	// Only allow GET requests.
	if r.Method != http.MethodGet {
		http.Error(w, "request method must be "+http.MethodGet, http.StatusMethodNotAllowed)
		return
	}

	// Path should be for the contributors endpoint.
	if !pathRE.MatchString(r.URL.Path) {
		http.Error(w, "url path should be /repos/{org}/{repo}/contributors", http.StatusNotFound)
		return
	}

	// Write json data to response.
	io.WriteString(w, contributors)
}

func main() {

	// Start listening for local traffic on port 0 which tells the OS to pick a
	// random open port. We start the listener seperately from the server so we
	// can report the listener's address.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mock GitHub server listening on", listener.Addr().String())
	log.Fatal(http.Serve(listener, http.HandlerFunc(handle)))
}
