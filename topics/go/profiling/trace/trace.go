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
	item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []item   `xml:"item"`
	}

	document struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	trace.Start(os.Stdout)
	defer trace.Stop()

	docs := make([]string, 100)
	for i := range docs {
		docs[i] = "newsfeed.xml"
	}

	topic := "president"
	n := findSingle(topic, docs)
	// n := findConcurrent(topic, docs)
	// n := findNumCPU(topic, docs)
	log.Printf("Found %s %d times.", topic, n)
}

func findSingle(topic string, docs []string) int {
	var found int

	for _, doc := range docs {
		f, err := os.Open(doc)
		if err != nil {
			log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
			return 0
		}
		defer f.Close()

		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
			return 0
		}

		var d document
		if err := xml.Unmarshal(data, &d); err != nil {
			log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
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

func findConcurrent(topic string, docs []string) int {
	var found int32

	var wg sync.WaitGroup
	wg.Add(len(docs))

	for _, doc := range docs {
		go func(doc string) {
			defer wg.Done()

			f, err := os.Open(doc)
			if err != nil {
				log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
				return
			}
			defer f.Close()

			data, err := ioutil.ReadAll(f)
			if err != nil {
				log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
				return
			}

			var d document
			if err := xml.Unmarshal(data, &d); err != nil {
				log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
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
		}(doc)
	}

	wg.Wait()
	return int(found)
}

func findNumCPU(topic string, docs []string) int {
	var found int32

	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	const workers = 8
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			var localFound int32

			for doc := range ch {
				f, err := os.Open(doc)
				if err != nil {
					log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
					return
				}
				defer f.Close()

				data, err := ioutil.ReadAll(f)
				if err != nil {
					log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
					return
				}

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
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
