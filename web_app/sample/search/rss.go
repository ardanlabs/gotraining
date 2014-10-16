// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package search : rss provides RSS feed support.
package search

import (
	"encoding/xml"
	"log"
	"net/http"
)

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

// rssSearch is used against any RSS feeds.
func rssSearch(engine string, uri string, searchResults chan<- []Result) {
	log.Printf("%s : rssSearch : Started : URI[%s]\n", engine, uri)

	// Slice for the results.
	var results []Result

	// On return send the results we have.
	defer func() {
		log.Printf("%s : rssSearch : Info : Sending Results\n", engine)
		searchResults <- results
	}()

	// Issue the search against the engine.
	resp, err := http.Get(uri)
	if err != nil {
		log.Printf("%s : rssSearch : Get : ERROR : %s\n", engine, err)
		return
	}

	// Schedule the close of the response body.
	defer resp.Body.Close()

	// Decode the results into the slice of maps.
	var d Document
	err = xml.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Printf("%s : rssSearch : Decode : ERROR : %s\n", engine, err)
		return
	}

	// Capture the data we need for our results.
	for _, result := range d.Channel.Items {
		results = append(results, Result{
			Engine:  engine,
			Title:   result.Title,
			Link:    result.Link,
			Content: result.Description,
		})
	}

	log.Printf("%s : rssSearch : Completed : Found[%d]\n", engine, len(results))
}
