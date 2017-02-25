package search

import (
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// Maintain a cache of retrieved documents. The cache will maintain items for
// fifteen seconds and then marked as expired. This is a very small cache so the
// gc time will be every hour.

const (
	expiration = time.Second * 15
	cleanup    = time.Hour
)

var cache = gc.New(expiration, cleanup)

// =============================================================================

type (
	// Item defines the fields associated with the item tag in the buoy RSS document.
	Item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
	}

	// Image defines the fields associated with the image tag in the buoy RSS document.
	Image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// Channel defines the fields associated with the channel tag in the buoy RSS document.
	Channel struct {
		XMLName xml.Name `xml:"channel"`
		Image   Image    `xml:"image"`
		Items   []Item   `xml:"item"`
	}

	// Document defines the fields associated with the buoy RSS document.
	Document struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
	}
)

// =============================================================================

// rssSearch is used against any RSS feeds.
func rssSearch(uid, term, engine, uri string) ([]Result, error) {
	log.Printf("%s : %s-%s : rssSearch : Started\n", uid, engine, uri)

	var d Document
	key := term + engine + uri

	// Look in the cache.
	v, found := cache.Get(key)

	// Based on the cache lookup determine what to do.
	switch {
	case found:
		log.Printf("%s : %s-%s : rssSearch : CACHE : Found\n", uid, engine, uri)
		d = v.(Document)

	default:
		log.Printf("%s : %s-%s : rssSearch : GET\n", uid, engine, uri)

		// Pull down the rss feed.
		resp, err := http.Get(uri)
		if err != nil {
			log.Printf("%s : %s-%s : rssSearch : ERROR : %s\n", uid, engine, uri, err)
			return []Result{}, err
		}

		// Schedule the close of the response body.
		defer resp.Body.Close()

		// Decode the results into a document.
		err = xml.NewDecoder(resp.Body).Decode(&d)
		if err != nil {
			log.Printf("%s : %s-%s : rssSearch : ERROR : Decode : %s\n", uid, engine, uri, err)
			return []Result{}, err
		}

		// Save this document into the cache.
		cache.Set(key, d, expiration)
	}

	// Create an empty slice of results.
	results := []Result{}

	// Capture the data we need for our results if we find the search term.
	for _, result := range d.Channel.Items {
		if strings.Contains(strings.ToLower(result.Description), term) {
			results = append(results, Result{
				Engine:  engine,
				Title:   result.Title,
				Link:    result.Link,
				Content: result.Description,
			})
		}
	}

	log.Printf("%s : %s-%s : rssSearch : Completed : Found[%d]\n", uid, engine, uri, len(results))
	return results, nil
}
