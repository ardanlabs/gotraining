## Embedding - Methods, Interfaces and Embedding

Embedding types provides the final piece of sharing and reusing state and behavior between types. Through the use of inner type promotion, an inner type's fields and methods can be directly access by references of the outer type.

## Notes

* Embedding types allow us to share state or behavior between types.
* The inner type never loses its identity.
* This is not inheritance.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* The outer type can override the inner type's behavior.

## Links

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## Code Review

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/Bweb5f-xdM))

[Embedding types](example2/example2.go) ([Go Playground](http://play.golang.org/p/wAV3xnKj60))

[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/_MiwwXZbVI))

[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/QSjaJocj5S))

## Exercises

### Exercise 1

Copy the code from the template. Declare a new type called hockey which embeds the sports type. Implement the matcher interface for hockey. When implementing the Search method for hockey, call into the Search method for the embedded sport type to check the embedded fields first. Then create two hockey values inside the slice of matchers and perform the search.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/c9Qrsq8QFe)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/B7VVYyaA21))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
