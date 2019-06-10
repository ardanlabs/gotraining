package search

import (
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// Maintain a cache of retrieved documents. The cache will maintain items for
// fifteen seconds and then marked as expired. This is a very small cache so the
// gc time will be every hour.

const (
	expiration = time.Minute * 15
	cleanup    = time.Hour
)

var cache = gc.New(expiration, cleanup)

var fetch = struct {
	sync.Mutex
	m map[string]*sync.Mutex
}{
	m: make(map[string]*sync.Mutex),
}

type (

	// Item defines the fields associated with the item tag in the RSS document.
	Item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
	}

	// Image defines the fields associated with the image tag in the RSS document.
	Image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// Channel defines the fields associated with the channel tag in the RSS document.
	Channel struct {
		XMLName xml.Name `xml:"channel"`
		Image   Image    `xml:"image"`
		Items   []Item   `xml:"item"`
	}

	// Document defines the fields associated with the RSS document.
	Document struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
	}
)

// rssSearch is used against any RSS feeds.
func rssSearch(uid, term, engine, uri string) ([]Result, error) {
	var mu *sync.Mutex
	fetch.Lock()
	{
		var found bool
		mu, found = fetch.m[uri]
		if !found {
			mu = &sync.Mutex{}
			fetch.m[uri] = mu
		}
	}
	fetch.Unlock()

	var d Document
	mu.Lock()
	{
		// Look in the cache.
		v, found := cache.Get(uri)

		// Based on the cache lookup determine what to do.
		switch {
		case found:
			d = v.(Document)

		default:

			// Pull down the rss feed.
			resp, err := http.Get(uri)
			if err != nil {
				return []Result{}, err
			}

			// Schedule the close of the response body.
			defer resp.Body.Close()

			// Decode the results into a document.
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				return []Result{}, err
			}

			// Save this document into the cache.
			cache.Set(uri, d, expiration)

			log.Println("reloaded cache", uri)
		}
	}
	mu.Unlock()

	// Create an empty slice of results.
	results := []Result{}

	// Capture the data we need for our results if we find the search term.
	for _, item := range d.Channel.Items {
		if strings.Contains(strings.ToLower(item.Description), strings.ToLower(term)) {
			results = append(results, Result{
				Engine:  engine,
				Title:   item.Title,
				Link:    item.Link,
				Content: item.Description,
			})
		}
	}

	return results, nil
}
