// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template2

// Sample program that commits data into pachyderm's data versioning.
package main

import (
	"log"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Start a commit in our "diabetes" data repo on the "master" branch.

	// Open the diabetes dataset file.

	// Put a file containing the diabetes data into the data repository.

	// Finish the commit.

	// As a sanity check, let's check to see that the commit
	// actually happened. The nil, 1, 0, and false parameters are
	// good defaults here and are explained further in the docs at
	// https://godoc.org/github.com/pachyderm/pachyderm/src/client#APIClient.ListCommitByRepo.

	// Check that the number of commits is what we expect.

	// Check that the ID of the commit is what we expect.
}
