package search

import "log"

var nytFeeds = []string{
	"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/US.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Politics.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Business.xml",
}

// NYT provides support for NYT searches.
type NYT struct{}

// NewNYT returns a NYT Searcher value.
func NewNYT() Searcher {
	return NYT{}
}

// Search performs a search against the NYT RSS feeds.
func (NYT) Search(uid string, term string, found chan<- []Result) {
	log.Printf("%s : NYT : Search : Started : searchTerm[%s]\n", uid, term)

	results := []Result{}

	for _, feed := range nytFeeds {
		res, err := rssSearch(uid, term, "NYT", feed)
		if err != nil {
			continue
		}

		results = append(results, res...)
	}

	found <- results

	log.Printf("%s : NYT : Search : Completed\n", uid)
}
