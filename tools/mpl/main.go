// Package main provides a CLI tool to automatically update Go Playground
// links within a training markdown file
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"regexp"
)

var (
	// rgxLink pulls out the title and source code file name.
	// [Title](example1/example1.go)
	rgxLink = regexp.MustCompile("\\[([^\\]]+)\\]\\(([^\\)]+)\\)")

	// rgxLinkInfo matches what source code files have a bound playground link.
	// [Title](example1/example1.go) ([Go Playground](http://play.golang.org/p/CoBIh_6Hjj))
	rgxLinkInfo = regexp.MustCompile("\\[[^\\]]+\\]\\([^\\)]+\\.go\\) +\\(\\[Go Playground\\]\\([^\\)]*\\)\\)")
)

// =============================================================================

func main() {
	const version = "1.1"

	// Parse the flags from the command line call.
	flag.Parse()

	// For sanity checks with multiple versions of the tool.
	if len(flag.Args()) == 0 {
		log.SetFlags(0)
		log.Println("Ver", version)
		return
	}

	// A list of files are expected to be passed on the command line.
	// If you run zsh you can use this support: mpl **/*.md

	for _, file := range flag.Args() {
		if err := process(file); err != nil {
			log.Fatal(err)
		}
	}
}

// process looks for source code links in a given markdown file. For every
// link that is found, the source code is run through the playground to generate
// a new playground link. Then the markdown file is updated.
func process(mdFile string) error {
	log.Println("Processing", mdFile)

	// Read in the entire markdown file.
	srcMd, err := ioutil.ReadFile(mdFile)
	if err != nil {
		return err
	}

	// This function will be used to process every source code link that
	// is found inside the markdown file based on the "rgxLinkInfo" regular exp.
	f := func(linkInfo []byte) []byte {

		// linkInfo : [Title](example1/example1.go) ([Go Playground](http://play.golang.org/p/CoBIh_6Hjj))

		// Create a match value to extract the title and source code file name.
		// [Title](example1/example1.go)
		m := rgxLink.FindAllSubmatch(linkInfo, 1)

		// Extract the title and source file name from the match.
		title := string(m[0][1])
		srcFile := string(m[0][2])

		// Read in the contents of the source code file.
		srcCode, err := ioutil.ReadFile(path.Join(path.Dir(mdFile), srcFile))
		if err != nil {
			log.Fatalf("ERROR: Reading Title[%s] ScrFile[%s] : %v", title, srcFile, err)
		}

		// Send the contents of the source code file to the playground.
		// Generate a new link for this code.
		playLink, err := generatePlayLink(srcCode)
		if err != nil {
			log.Fatalf("ERROR: Generating Link Title[%s] ScrFile[%s] : %v", title, srcFile, err)
		}

		// Generate the link information to replace the existing link infomation
		// in the markdown file.
		l := fmt.Sprintf("[%s](%s) ([Go Playground](%s))", title, srcFile, playLink)
		log.Println("UPDATE:", l)

		return []byte(l)
	}

	// Find all matches for the link info regex against the markdown file.
	// For every match, execute the anonymous function above to replace the
	// current playground link with an updated one inside the markdown file.
	res := rgxLinkInfo.ReplaceAllFunc(srcMd, f)

	// Write the new markdown file back to disk with proper permissions.
	return ioutil.WriteFile(mdFile, res, 077)
}

// generatePlayLink returns the URL for the playground link based on the
// source code that is provided.
func generatePlayLink(srcCode []byte) (string, error) {
	const url = "http://play.golang.org/share"
	const mime = "application/x-www-form-urlencoded; charset=UTF-8"

	// Make a call to the playground, posting the source code file contents.
	res, err := http.Post(url, mime, bytes.NewReader(srcCode))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected playground status code: %d", res.StatusCode)
	}

	// Read back the generated URL GUID for this source code.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://play.golang.org/p/%s", string(b)), nil
}
