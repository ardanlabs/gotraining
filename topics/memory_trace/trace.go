// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to learn how to identify allocations and work through them.
package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
)

// Item defines the fields associated with the item tag in
// the buoy RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in
// the buoy RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

var doc = make(map[int]*Document)

func main() {

	go func() {
		var id int

		for {
			id++
			d, err := get(id)
			if err != nil {
				continue
			}

			fmt.Println(d.Channel.Title)

			time.Sleep(time.Millisecond * 25)
		}
	}()

	// Start a listener for the pprof support.
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	// On a <ctrl> C shutdown the program.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	<-sig
}

func get(id int) (*Document, error) {
	const url = "https://www.goinggo.net/post/index.xml"

	if d, exists := doc[id]; exists {
		return d, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var d Document
	if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}

	doc[id] = &d

	return &d, nil
}
