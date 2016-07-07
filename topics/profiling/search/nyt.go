// Package search : bing performs searches against the bing search engine.
package search

import "log"

var nytFeeds = []string{
	"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/US.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Politics.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Business.xml",
}

// NYTimes provides support for CNN searches.
type NYTimes struct{}

// NewNYTimes returns a NYTimes Searcher value.
func NewNYTimes() Searcher {
	return NYTimes{}
}

// Search performs a search against the CNN RSS feeds.
func (NYTimes) Search(term string, found chan<- []Result) {
	log.Printf("NYTimes : Search : Started : searchTerm[%s]\n", term)

	results := []Result{}

	for _, feed := range nytFeeds {
		res, err := rssSearch(term, "NYTimes", feed)
		if err != nil {
			continue
		}

		results = append(results, res...)
	}

	found <- results

	log.Println("NYTimes : Search : Completed")
}
