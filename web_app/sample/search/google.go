// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package search : google performs searches against the google search engine.
package search

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Google provides support for Google searches.
type Google struct{}

type (
	// gResult maps to the result document received from the search.
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// gResponse contains the top level document.
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

// NewGoogle returns a Google Searcher value.
func NewGoogle() Searcher {
	return Google{}
}

// Search implements the Searcher interface. It performs a search
// against Google.
func (g Google) Search(searchTerm string, searchResults chan<- []Result) {
	log.Printf("Google : Search : Started : searchTerm[%s]\n", searchTerm)

	// Slice for the results.
	var results []Result

	// On return send the results we have.
	defer func() {
		log.Println("Google : Search : Info : Sending Results")
		searchResults <- results
	}()

	// Build a proper search url.
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=" + searchTerm

	// Issue the search against Google.
	resp, err := http.Get(uri)
	if err != nil {
		log.Printf("Google : Search : Get : ERROR : %s\n", err)
		return
	}

	// Schedule the close of the response body.
	defer resp.Body.Close()

	// Decode the JSON response into our struct type.
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Printf("Google : Search : Decode : ERROR : %s\n", err)
		return
	}

	// Capture the data we need for our results.
	for _, result := range gr.ResponseData.Results {
		results = append(results, Result{
			Engine:  "Google",
			Title:   result.Title,
			Link:    result.URL,
			Content: result.Content,
		})
	}

	log.Printf("Google : Search : Completed : Found[%d]\n", len(results))
}
