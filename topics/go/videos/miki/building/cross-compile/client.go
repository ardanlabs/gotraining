package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: quotes")
		flag.PrintDefaults()
	}
	flag.Parse()

	resp, err := http.Get("https://quotamiki.appspot.com/quote")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: %s", err)
	}
}
