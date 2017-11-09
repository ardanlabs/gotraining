// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
)

type (

	// Item defines the fields associated with the item tag in the buoy RSS document.
	Item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}

	// Channel defines the fields associated with the channel tag in the buoy RSS document.
	Channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []Item   `xml:"item"`
	}

	// Document defines the fields associated with the buoy RSS document.
	Document struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
	}
)

func findSingle(topic string) int {

	// Track the number of articles that contain
	// the topic.
	var found int

	// Pretend we have 100 feeds to search.
	for i := 0; i < 100; i++ {
		feed := "newsfeed.xml"

		// Open the file to search.
		f, err := os.Open(feed)
		if err != nil {
			log.Printf("Opening File [%s] : ERROR : %v", feed, err)
			return 0
		}
		defer f.Close()

		// Read the string into memory.
		doc, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("Reading File [%s] : ERROR : %v", feed, err)
			return 0
		}

		// Decode the RSS document.
		var d Document
		if err := xml.Unmarshal(doc, &d); err != nil {
			log.Printf("Decoding File [%s] : ERROR : %v", feed, err)
			f.Close()
			return 0
		}

		// Find the topic.
		for _, item := range d.Channel.Items {
			if strings.Contains(item.Title, topic) {
				found++
				continue
			}

			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}

	return found
}

func findConcurrent(topic string) int {

	// Track the number of articles that contain
	// the topic.
	var found int32

	var wg sync.WaitGroup
	wg.Add(100)

	// Pretend we have 100 feeds to search.
	for i := 0; i < 100; i++ {
		go func() {
			feed := "newsfeed.xml"

			// Open the file to search.
			f, err := os.Open(feed)
			if err != nil {
				log.Printf("Opening File [%s] : ERROR : %v", feed, err)
				return
			}
			defer f.Close()

			// Read the string into memory.
			doc, err := ioutil.ReadAll(f)
			if err != nil {
				log.Printf("Reading File [%s] : ERROR : %v", feed, err)
				return
			}

			// Decode the RSS document.
			var d Document
			if err := xml.Unmarshal(doc, &d); err != nil {
				log.Printf("Decoding File [%s] : ERROR : %v", feed, err)
				f.Close()
				return
			}

			var localFound int32

			// Find the topic.
			for _, item := range d.Channel.Items {
				if strings.Contains(item.Title, topic) {
					localFound++
					continue
				}

				if strings.Contains(item.Description, topic) {
					localFound++
				}
			}

			atomic.AddInt32(&found, localFound)
			wg.Done()
		}()
	}

	wg.Wait()

	return int(found)
}

func findLimit(topic string) int {

	// Track the number of articles that contain
	// the topic.
	var found int32

	var wg sync.WaitGroup
	wg.Add(8)

	ch := make(chan bool, 100)

	// Pretend we have 100 feeds to search.
	for i := 0; i < 8; i++ {
		go func() {
			var localFound int32

			for range ch {
				feed := "newsfeed.xml"

				// Open the file to search.
				f, err := os.Open(feed)
				if err != nil {
					log.Printf("Opening File [%s] : ERROR : %v", feed, err)
					return
				}
				defer f.Close()

				// Read the string into memory.
				doc, err := ioutil.ReadAll(f)
				if err != nil {
					log.Printf("Reading File [%s] : ERROR : %v", feed, err)
					return
				}

				// Decode the RSS document.
				var d Document
				if err := xml.Unmarshal(doc, &d); err != nil {
					log.Printf("Decoding File [%s] : ERROR : %v", feed, err)
					f.Close()
					return
				}

				// Find the topic.
				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						localFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						localFound++
					}
				}
			}

			atomic.AddInt32(&found, localFound)
			wg.Done()
		}()
	}

	for i := 0; i < 100; i++ {
		ch <- true
	}
	close(ch)

	wg.Wait()

	return int(found)
}

func main() {

	// Start gathering the tracing data.
	trace.Start(os.Stdout)
	defer trace.Stop()

	topic := "president"
	n := findSingle(topic)
	log.Printf("Found %s %d times.", topic, n)
}
