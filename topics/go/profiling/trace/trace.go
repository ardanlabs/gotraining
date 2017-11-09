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
	"runtime/pprof"
	"strings"
	"sync"
	"sync/atomic"
)

type (
	// Item defines the fields associated with the item tag in the RSS document.
	Item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}

	// Channel defines the fields associated with the channel tag in the RSS document.
	Channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []Item   `xml:"item"`
	}

	// Document defines the fields associated with the RSS document.
	Document struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
	}
)

func findSingle(topic string) int {
	const feeds = 100

	var found int

	for i := 0; i < feeds; i++ {
		feed := "newsfeed.xml"
		f, err := os.Open(feed)
		if err != nil {
			log.Printf("Opening File [%s] : ERROR : %v", feed, err)
			return 0
		}
		defer f.Close()

		doc, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("Reading File [%s] : ERROR : %v", feed, err)
			return 0
		}

		var d Document
		if err := xml.Unmarshal(doc, &d); err != nil {
			log.Printf("Decoding File [%s] : ERROR : %v", feed, err)
			f.Close()
			return 0
		}

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
	const feeds = 100

	var found int32

	var wg sync.WaitGroup
	wg.Add(feeds)

	for i := 0; i < feeds; i++ {
		go func() {
			defer wg.Done()

			feed := "newsfeed.xml"
			f, err := os.Open(feed)
			if err != nil {
				log.Printf("Opening File [%s] : ERROR : %v", feed, err)
				return
			}
			defer f.Close()

			doc, err := ioutil.ReadAll(f)
			if err != nil {
				log.Printf("Reading File [%s] : ERROR : %v", feed, err)
				return
			}

			var d Document
			if err := xml.Unmarshal(doc, &d); err != nil {
				log.Printf("Decoding File [%s] : ERROR : %v", feed, err)
				f.Close()
				return
			}

			var localFound int32
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
		}()
	}

	wg.Wait()
	return int(found)
}

func findWorkers(topic string) int {
	const feeds = 100
	const workers = 8

	var found int32

	ch := make(chan string, feeds)
	for i := 0; i < feeds; i++ {
		ch <- "newsfeed.xml"
	}
	close(ch)

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			var localFound int32

			for feed := range ch {
				f, err := os.Open(feed)
				if err != nil {
					log.Printf("Opening File [%s] : ERROR : %v", feed, err)
					return
				}
				defer f.Close()

				doc, err := ioutil.ReadAll(f)
				if err != nil {
					log.Printf("Reading File [%s] : ERROR : %v", feed, err)
					return
				}

				var d Document
				if err := xml.Unmarshal(doc, &d); err != nil {
					log.Printf("Decoding File [%s] : ERROR : %v", feed, err)
					f.Close()
					return
				}

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
		}()
	}

	wg.Wait()
	return int(found)
}

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// trace.Start(os.Stdout)
	// defer trace.Stop()

	topic := "president"
	n := findSingle(topic)
	log.Printf("Found %s %d times.", topic, n)
}
