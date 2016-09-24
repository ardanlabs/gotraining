## Embedding

Embedding types provides the final piece of sharing and reusing state and behavior between types. Through the use of inner type promotion, an inner type's fields and methods can be directly accessed by references of the outer type.

## Notes

* Embedding types allows us to share state or behavior between types.
* The inner type never loses its identity.
* This is not inheritance.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* The outer type can override the inner type's behavior.

## Links

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## Code Review

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/P3Ynb7eqBL))  
[Embedding types](example2/example2.go) ([Go Playground](https://play.golang.org/p/QncBd6A5A4))  
[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/vMEEJ7rOb4))  
[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/je8_2_Wny-))

## Exercises

### Exercise 1

Copy the code from the template. Declare a new type called hockey which embeds the sports type. Implement the matcher interface for hockey. When implementing the match method for hockey, call into the match method for the embedded sport type to check the embedded fields first. Then create two hockey values inside the slice of matchers and perform the search.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/e8Yz-hmypo)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/v6vjpkyBdN))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
