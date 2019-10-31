// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program defines a type Feed with two methods: Count and Fetch. Create a
// new type CachingFeed that embeds *Feed but overrides the Fetch method.
//
// The CachingFeed type should have a map of Documents to limit the number of
// calls to Feed.Fetch.
package main

import (
	"fmt"
	"log"
	"time"
)

// Document is the core data model we are working with.
type Document struct {
	Key   string
	Title string
}

// ==================================================

// Feed is a type that knows how to fetch Documents.
type Feed struct{}

// Count tells how many documents are in the feed.
func (f *Feed) Count() int {
	return 42
}

// Fetch simulates looking up the document specified by key. It is slow.
func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

// ==================================================

// FetchCounter is the behavior we depend on for our process function.
type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}

		fmt.Printf("%s : %v\n", key, doc)
	}
}

// ==================================================

// CachingFeed keeps a local copy of Documents that have already been
// retrieved. It embeds Feed to get the Fetch and Count behavior but
// "overrides" Fetch to have its cache.
type CachingFeed struct {
	// TODO embed *Feed and add a field for a map[string]Document.
}

// NewCachingFeed initializes a CachingFeed for use.
func NewCachingFeed(f *Feed) *CachingFeed {

	// TODO create a CachingFeed with an initialized map and embedded feed.
	// Return its address.

	return nil // Remove this line.
}

// Fetch calls the embedded type's Fetch method if the key is not cached.
func (cf *CachingFeed) Fetch(key string) (Document, error) {

	// TODO implement this method. Check the map field for the specified key and
	// return it if found. If it's not found, call the embedded types Fetch
	// method. Store the result in the map before returning it.

	return Document{}, nil // Remove this line.
}

// ==================================================

func main() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	// Call process again with your CachingFeed.
	//fmt.Println("Using CachingFeed")
}
