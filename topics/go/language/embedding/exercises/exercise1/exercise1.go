// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how you can use embedding to reuse behavior from
// another type and override specific methods.
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

// CachingFeed keeps a local copy of Documents that have already been
// retrieved. It embeds Feed to get the Fetch and Count behavior but
// "overrides" Fetch to have its cache.
type CachingFeed struct {
	docs map[string]Document
	*Feed
}

// NewCachingFeed initializes a CachingFeed for use.
func NewCachingFeed(f *Feed) *CachingFeed {
	return &CachingFeed{
		docs: make(map[string]Document),
		Feed: f,
	}
}

// Fetch calls the embedded type's Fetch method if the key is not cached.
func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, ok := cf.docs[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.docs[key] = doc
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

func main() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	fmt.Println("Using CachingFeed")
	c := NewCachingFeed(&Feed{})
	process(c)
}
