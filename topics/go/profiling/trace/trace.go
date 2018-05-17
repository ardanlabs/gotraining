// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// RUN: ulimit -n 20000

// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"runtime"
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

	// trace.Start(os.Stdout)
	// defer trace.Stop()

	docs := make([]string, 20000)
	for i := range docs {
		docs[i] = "newsfeed.xml"
	}

	topic := "president"
	n := find(topic, docs)
	// n := findConcurrent(topic, docs)
	// n := findConcurrentSem(topic, docs)
	// n := findNumCPU(topic, docs)
	// n := findActor(topic, docs)

	log.Printf("Found %s %d times.", topic, n)
}

func find(topic string, docs []string) int {
	var found int

	for _, doc := range docs {
		f, err := os.OpenFile(doc, os.O_RDONLY, 0)
		if err != nil {
			log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
			return 0
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			f.Close()
			log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
			return 0
		}
		f.Close()

		var d document
		if err := xml.Unmarshal(data, &d); err != nil {
			log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
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

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	for _, doc := range docs {
		go func(doc string) {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			f, err := os.OpenFile(doc, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
				return
			}

			data, err := ioutil.ReadAll(f)
			if err != nil {
				f.Close()
				log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
				return
			}
			f.Close()

			var d document
			if err := xml.Unmarshal(data, &d); err != nil {
				log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
				return
			}

			for _, item := range d.Channel.Items {
				if strings.Contains(item.Title, topic) {
					lFound++
					continue
				}

				if strings.Contains(item.Description, topic) {
					lFound++
				}
			}
		}(doc)
	}

	wg.Wait()
	return int(found)
}

func findConcurrentSem(topic string, docs []string) int {
	var found int32

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan bool, runtime.NumCPU())

	for _, doc := range docs {
		go func(doc string) {
			ch <- true
			{
				var lFound int32
				defer func() {
					atomic.AddInt32(&found, lFound)
					wg.Done()
				}()

				f, err := os.OpenFile(doc, os.O_RDONLY, 0)
				if err != nil {
					log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
					return
				}

				data, err := ioutil.ReadAll(f)
				if err != nil {
					f.Close()
					log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
					return
				}
				f.Close()

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
					return
				}

				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
			<-ch
		}(doc)
	}

	wg.Wait()
	return int(found)
}

func findNumCPU(topic string, docs []string) int {
	var found int32

	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	for i := 0; i < g; i++ {
		go func() {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			for doc := range ch {
				f, err := os.OpenFile(doc, os.O_RDONLY, 0)
				if err != nil {
					log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
					return
				}

				data, err := ioutil.ReadAll(f)
				if err != nil {
					f.Close()
					log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
					return
				}
				f.Close()

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
					return
				}

				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
		}()
	}

	wg.Wait()
	return int(found)
}

func findActor(topic string, docs []string) int {
	files := make(chan *os.File, 100)
	go func() {
		for _, doc := range docs {
			f, err := os.OpenFile(doc, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
				break
			}
			files <- f
		}
		close(files)
	}()

	data := make(chan []byte, 100)
	go func() {
		for f := range files {
			defer f.Close()
			d, err := ioutil.ReadAll(f)
			if err != nil {
				log.Printf("Reading Document [%s] : ERROR : %v", f.Name(), err)
				break
			}
			data <- d
		}
		close(data)
	}()

	rss := make(chan document, 100)
	go func() {
		for dt := range data {
			var d document
			if err := xml.Unmarshal(dt, &d); err != nil {
				log.Printf("Decoding Document : ERROR : %v", err)
				break
			}
			rss <- d
		}
		close(rss)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	var found int
	go func() {
		for d := range rss {
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
		wg.Done()
	}()

	wg.Wait()
	return found
}
