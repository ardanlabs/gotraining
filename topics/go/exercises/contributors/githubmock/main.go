package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
)

var pathRE = regexp.MustCompile("^/repos/([^/]+/[^/]+)/contributors$")
var authRE = regexp.MustCompile("^token .+$")

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.Method, r.URL.Path)

	// They must provide an auth token.
	if !authRE.MatchString(r.Header.Get("Authorization")) {
		http.Error(w, "Authorization header must be in the form \"token {githubToken}\"", http.StatusUnauthorized)
		return
	}

	// Only allow GET requests.
	if r.Method != "GET" {
		http.Error(w, "request method must be GET", http.StatusMethodNotAllowed)
		return
	}

	// Path should be for the contributors endpoint.
	pathMatches := pathRE.FindStringSubmatch(r.URL.Path)
	if len(pathMatches) == 0 {
		http.Error(w, "url path should be /repos/{org}/{repo}/contributors", http.StatusNotFound)
		return
	}

	// Look up contributors based on the path they specified
	var resp string
	switch pathMatches[1] {
	case "ardanlabs/gotraining":
		resp = ardanContributors
	case "golang/go":
		resp = goContributors
	default:
		http.Error(w, "unknown repo "+pathMatches[1], http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write json data to response.
	io.WriteString(w, resp)
}

func main() {

	// Start listening for local traffic on port 0 which tells the OS to pick a
	// random open port. We start the listener separately from the server so we
	// can report the listener's address.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mock GitHub server running at this url")
	fmt.Println("http://" + listener.Addr().String() + "/repos/golang/go/contributors")
	log.Fatal(http.Serve(listener, http.HandlerFunc(handle)))
}
