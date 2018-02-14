package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/gotraining/topics/go/exercises/contributors/part2/github"
)

func main() {
	tkn := os.Getenv("GITHUB_TOKEN")
	if tkn == "" {
		log.Print("Token not found. You must set it in your environment like")
		log.Print("export GITHUB_TOKEN=000a0aaaa0000a00000000aaa00000000a000000")
		log.Print("You can generate a token at https://github.com/settings/tokens")
		os.Exit(1)
	}

	c, err := github.NewClient(tkn)
	if err != nil {
		log.Fatal(err)
	}

	result := printContributors("ardanlabs/gotraining", c)

	os.Exit(result)
}

// contributors is the interface that this package looks for when
// calling printContributors.
type contributors interface {
	Contributors(string) ([]github.Contributor, error)
}

func printContributors(repo string, c contributors) int {
	cons, err := c.Contributors(repo)
	if err != nil {
		fmt.Printf("Error fetching contributors: %v\n", err)
		return 1
	}

	for i, con := range cons {
		fmt.Println(i, con.Login, con.Contributions)
	}

	return 0
}
