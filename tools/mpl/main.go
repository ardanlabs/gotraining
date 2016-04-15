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
	// linkRegexp matches markdown links.
	linkRegexp = regexp.MustCompile("\\[([^\\]]+)\\]\\(([^\\)]+)\\)")

	// srcLinkRegexp matches linked source file that has to be uploaded to the
	// Go Playground.
	srcLinkRegexp = regexp.MustCompile("\\[[^\\]]+\\]\\([^\\)]+\\.go\\) +\\(\\[Go Playground\\]\\([^\\)]*\\)\\)")
)

// link represents data required to create a link.
type link struct {
	title    string
	srcFile  string
	playLink string
	src      []byte
}

// newLink returns a new link based on a source file.
func newLink(mdFile, title, srcFile string) (link, error) {
	l := link{
		title:   title,
		srcFile: srcFile,
	}

	var err error
	if l.src, err = ioutil.ReadFile(path.Join(path.Dir(mdFile), l.srcFile)); err != nil {
		return link{}, err
	}

	if l.playLink, err = l.generatePlayLink(); err != nil {
		return link{}, err
	}

	return l, nil
}

// String returns the new content with the generated playground link
func (l link) String() string {
	return fmt.Sprintf("[%s](%s) ([Go Playground](%s))", l.title, l.srcFile, l.playLink)
}

// generatePlayLink returns the URL to the playground for the linked source code file
func (l link) generatePlayLink() (string, error) {
	res, err := http.Post("http://play.golang.org/share", "application/x-www-form-urlencoded; charset=UTF-8", bytes.NewReader(l.src))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("unexpected playground status code: %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://play.golang.org/p/%s", string(b)), nil
}

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		if err := process(file); err != nil {
			log.Fatal(err)
		}
	}
}

// process looks for playground links in a given markdown files and
// generates new links.
func process(file string) error {
	log.Println("Processing", file)

	src, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	f := func(b []byte) []byte {
		m := linkRegexp.FindAllSubmatch(b, 1)

		title := string(m[0][1])
		srcFile := string(m[0][2])

		log.Println("Updating", title, srcFile)

		l, err := newLink(file, title, srcFile)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(l)
		return []byte(l.String())
	}

	res := srcLinkRegexp.ReplaceAllFunc(src, f)

	return ioutil.WriteFile(file, res, 077)
}
