// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// You are going to want to copy and paste this line later.
// "github.com/ardanlabs/gotraining/topics/go/exercises/contributors/template/github"

// Call the GitHub API to get a list of repository contributors.
package main

import (
	"log"
	"os"
)

// Create a type where we can decode contributor json values.
// It needs the fields "login" and "contributions".

func main() {

	// Get an access token from the environment.
	tkn := os.Getenv("GITHUB_TOKEN")
	if tkn == "" {
		log.Print("Token not found. You must set it in your environment like")
		log.Print("export GITHUB_TOKEN=000a0aaaa0000a00000000aaa00000000a000000")
		log.Print("You can generate a token at https://github.com/settings/tokens")
		os.Exit(1)
	}

	// Create a request for the contributors api endpoint.

	// Add the access token in the "Authorization" header.
	// The value should be in the form "token 000aa0a0..."

	// Create an http.Client and make the request.

	// Defer closing the response body.

	// Ensure we get a 200 OK status back.

	// Decode the results into a []contributor.

	// Loop through the []contributor and print the values.
}
