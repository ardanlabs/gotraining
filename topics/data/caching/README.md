## Caching

As already discussed, the day-to-day activities of a data scientists usually involve a lot of gathering, loading, and organizing data.  In performing these duties, you will likely encounter pieces of data that you need frequently, are difficult to obtain quickly, or are otherwise nice to have close to your application.  There are a variety of great caching and embedded database options when working with Go.  We will explore a couple here, `go-cache` and `boltdb`.

## Notes

* A cache isn’t meant to be used as a persistent datastore.
* A cache created with `go-cache` can be safely used by multiple goroutines.

## Links

[go-cache Docs](https://patrickmn.com/projects/go-cache/)    
[Intro to BoltDB](https://npf.io/2014/07/intro-to-boltdb-painless-performant-persistence/)

## Code Review

[Cache data in memory](example1/example1.go)  
[Save data in an embedded key/value store](example2/example2.go)   

## Exercises

### Exercise 1

Cache the Citibike status codes in memory.  Then, GET the current Citibike station statuses, and output all stations that are "Not in Service."

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go) 

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
