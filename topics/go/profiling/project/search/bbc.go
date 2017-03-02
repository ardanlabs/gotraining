package search

import "log"

var bbcFeeds = []string{
	"http://feeds.bbci.co.uk/news/rss.xml",
	"http://feeds.bbci.co.uk/news/world/rss.xml",
	"http://feeds.bbci.co.uk/news/politics/rss.xml",
	"http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml",
}

// BBC provides support for BBC searches.
type BBC struct{}

// NewBBC returns a BBC Searcher value.
func NewBBC() Searcher {
	return BBC{}
}

// Search performs a search against the CNN RSS feeds.
func (BBC) Search(uid string, term string, found chan<- []Result) {
	results := []Result{}

	for _, feed := range bbcFeeds {
		res, err := rssSearch(uid, term, "BBC", feed)
		if err != nil {
			log.Println("ERROR: ", err)
			continue
		}

		results = append(results, res...)
	}

	found <- results
}
