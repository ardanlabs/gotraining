## Maps

Maps provide a data structure that allow for the storage and management of key/value pair data.

## Notes

* Maps provide a way to store and retrieve key/value pairs.
* The map key must be a value that can be used in an assignment statement.
* Iterating over a map is always random.

## Links

http://blog.golang.org/go-maps-in-action  
http://www.goinggo.net/2013/12/macro-view-of-map-internals-in-go.html  
[Keith Randall - Inside the Map Implementation](https://www.youtube.com/watch?v=Tl7mi9QmLns)

## Code Review

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](https://play.golang.org/p/EHfkoipKYF))  
[Map literals and delete](example2/example2.go) ([Go Playground](https://play.golang.org/p/B2klwmqmPZ))  
[Map key restrictions](example3/example3.go) ([Go Playground](https://play.golang.org/p/LZRHA7FG6s))  

## Exercises

### Exercise 1

Declare and make a map of integer values with a string as the key. Populate the map with five values and iterate over the map to display the key/value pairs.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/E2VFcOY1o6)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/uT_pwbOgNc))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
