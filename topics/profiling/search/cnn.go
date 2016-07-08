package search

import "log"

var cnnFeeds = []string{
	"http://rss.cnn.com/rss/cnn_topstories.rss",
	"http://rss.cnn.com/rss/cnn_world.rss",
	"http://rss.cnn.com/rss/cnn_us.rss",
	"http://rss.cnn.com/rss/cnn_allpolitics.rss",
}

// CNN provides support for CNN searches.
type CNN struct{}

// NewCNN returns a CNN Searcher value.
func NewCNN() Searcher {
	return CNN{}
}

// Search performs a search against the CNN RSS feeds.
func (CNN) Search(uid string, term string, found chan<- []Result) {
	log.Printf("%s : CNN : Search : Started : searchTerm[%s]\n", uid, term)

	results := []Result{}

	for _, feed := range cnnFeeds {
		res, err := rssSearch(uid, term, "CNN", feed)
		if err != nil {
			continue
		}

		results = append(results, res...)
	}

	found <- results

	log.Printf("%s : CNN : Search : Completed\n", uid)
}
