## Maps

Maps provide a data structure that allow for the storage and management of key/value pair data.

## Notes

* Maps provide a way to store and retrieve key/value pairs.
* Reading an absent key returns the zero value for the map's value type.
* Iterating over a map is always random.
* The map key must be a value that is comparable.
* Elements in a map are not addressable.
* Maps are a reference type.

## Links

[Go maps in action](https://blog.golang.org/go-maps-in-action) - Andrew Gerrand    
[Macro View of Map Internals In Go](https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html) - William Kennedy    
[Inside the Map Implementation](https://www.youtube.com/watch?v=Tl7mi9QmLns) - Keith Randall    
[How the Go runtime implements maps efficiently (without generics)](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics) - Dave Cheney    

## Code Review

[Declare, write, read, and delete](example1/example1.go) ([Go Playground](https://play.golang.org/p/3w6zgywPD3w))  
[Absent keys](example2/example2.go) ([Go Playground](https://play.golang.org/p/5KHMfmL2SyA))  
[Map key restrictions](example3/example3.go) ([Go Playground](https://play.golang.org/p/lfl967ocaKv))  
[Map literals and range](example4/example4.go) ([Go Playground](https://play.golang.org/p/0KFlxby2a0z))  
[Sorting maps by key](example5/example5.go) ([Go Playground](https://play.golang.org/p/XADXCQqn2pJ))  
[Taking an element's address](example6/example6.go) ([Go Playground](https://play.golang.org/p/4phv1S1wZWh))  
[Maps are Reference Types](example7/example7.go) ([Go Playground](https://play.golang.org/p/7jEDn1yhg5v))  

## Exercises

### Exercise 1

Declare and make a map of integer values with a string as the key. Populate the map with five values and iterate over the map to display the key/value pairs.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/FjQuvFWPz6m)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/KErzw53nM8A))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
