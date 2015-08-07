## Composition

Composition goes beyond the mechanics of type embedding. It's a design pattern we can leverage in Go to build larger programs from smaller parts. These smaller parts come from the declaration and implementation of types that have a single focus. Programs that are architected with composition in mind have a better chance to grow and adapt to changes.

## Notes

* Declare types and behavior with composition in mind.
* Composition is like building software with lego blocks.
* This is much more than the mechanics of type embedding.

## Links

http://golang.org/doc/effective_go.html#embedding

## Code Review

[Composition I](example1/example1.go) ([Go Playground](http://play.golang.org/p/W5ya6_LAU6))

[Composition II](example2/example2.go) ([Go Playground](http://play.golang.org/p/xsDJhCYOBA))

## Exercises

### Exercise 1

**Part A** Follow the guided comments to:

**Part B** Declare a sysadmin type that implements the administrator interface.

**Part C** Declare a programmer type that implements the developer interface.

**Part D** Declare a company type that embeds both an administrator and a developer.

**Part E** Create a sysadmin, programmers, and a company which are available for hire, and use them to complete some predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/b8ww3jd2Xs)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/8t5ns3cqNp))
___
All material is licensed under the [GNU Free Documentation License](https://github.com/gobridge/gotraining/blob/master/LICENSE).
