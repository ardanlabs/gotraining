// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program that commits data into Pachyderm data versioning.
package main

import (
	"log"
	"os"

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

	// Start a commit in our "iris" data repo on the "master" branch.
	commit, err := c.StartCommit("iris", "master")
	if err != nil {
		log.Fatal(err)
	}

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Put a file containing the iris data into the data repository.
	if _, err := c.PutFile("iris", commit.ID, "iris.csv", f); err != nil {
		log.Fatal(err)
	}

	// Finish the commit.
	if err := c.FinishCommit("iris", commit.ID); err != nil {
		log.Fatal(err)
	}

	// As a sanity check, let's check to see that the commit
	// actually happened. The nil, 1, 0, and false parameters are
	// good defaults here and are explained further in the docs at
	// https://godoc.org/github.com/pachyderm/pachyderm/src/client#APIClient.ListCommitByRepo.
	commits, err := c.ListCommitByRepo([]string{"iris"}, nil, 1, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	// Check that the number of commits is what we expect.
	if len(commits) != 1 {
		log.Fatal("Unexpected number of commits")
	}

	// Check that the ID of the commit is what we expect.
	if commits[0].Commit.ID != commit.ID {
		log.Fatal("Unexpected commit ID")
	}
}
