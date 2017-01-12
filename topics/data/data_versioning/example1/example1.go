// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program that connects to a running instance of Pachyderm.
package main

import (
	"fmt"
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

	// As a sanity check, we will list all the current data
	// repositories on the Pachyderm cluster.  Because we
	// haven't put anything in Pachyderm yet or created any
	// repositories, the number of these repos should be zero.
	repos, err := c.ListRepo(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Output the number of repos to standard out.
	fmt.Println(len(repos))
}
