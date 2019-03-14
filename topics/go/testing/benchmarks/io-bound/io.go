// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Understanding your workload is critically important in undserstanding
// if something can be made concurrent and how complex it is to perform.
package main

import (
	"encoding/xml"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	docs := generateList(1e3)

	fmt.Println(find("Go", docs))
	fmt.Println(findConcurrent(runtime.NumCPU(), "Go", docs))
}

func generateList(totalDocs int) []string {
	docs := make([]string, totalDocs)
	for i := 0; i < totalDocs; i++ {
		docs[i] = "test.xml"
	}
	return docs
}

func read(doc string) ([]item, error) {
	time.Sleep(time.Millisecond) // Simulate blocking disk read.
	var d document
	if err := xml.Unmarshal([]byte(file), &d); err != nil {
		return nil, err
	}
	return d.Channel.Items, nil
}

func find(topic string, docs []string) int {
	var found int
	for _, doc := range docs {
		items, err := read(doc)
		if err != nil {
			continue
		}
		for _, item := range items {
			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}
	return found
}

func findConcurrent(goroutines int, topic string, docs []string) int {
	var found int64

	var wg sync.WaitGroup
	wg.Add(goroutines)

	ch := make(chan string, len(docs))

	for g := 0; g < goroutines; g++ {
		go func() {
			var lFound int64
			for doc := range ch {
				items, err := read(doc)
				if err != nil {
					continue
				}
				for _, item := range items {
					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
			atomic.AddInt64(&found, lFound)
			wg.Done()
		}()
	}

	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	wg.Wait()
	return int(found)
}

var file = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an amazing language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

type item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []item   `xml:"item"`
}

type document struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
	URI     string
}
