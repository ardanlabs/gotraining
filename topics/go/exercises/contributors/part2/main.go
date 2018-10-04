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

	if err := process("ardanlabs/gotraining", c); err != nil {
		log.Fatal(err)
	}
}

// contributors is the interface that this package looks for when
// calling process.
type contributorLister interface {
	ContributorList(string) ([]github.Contributor, error)
}

func process(repo string, c contributorLister) error {
	cons, err := c.ContributorList(repo)
	if err != nil {
		return err
	}

	for i, con := range cons {
		fmt.Println(i, con.Login, con.Contributions)
	}

	return nil
}
