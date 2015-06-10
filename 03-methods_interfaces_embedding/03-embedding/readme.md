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

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/e5O_Dx5VpM))

[Embedding types](example2/example2.go) ([Go Playground](https://play.golang.org/p/UkrDXkk-Ch))

[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/BgEoThS7u9))

[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/jfOfrRMPZR))

[Composition](example5/example5.go) ([Go Playground](http://play.golang.org/p/QKIjbBDW16))

## Exercises

### Exercise 1

**Part A** Follow the guided comments to:

**Part B** Declare a sysadmin type that implements the administrator interface.

**Part C** Declare a programmer type that implements the developer interface.

**Part D** Declare a company type that embeds both an administrator and a developer.

**Part E** Create a sysadmin, programmers, and a company which are available for hire, and use them to complete some predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/zyXy90YfeY)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/hVFMZSUGI6))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
