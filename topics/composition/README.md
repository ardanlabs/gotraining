## Interface and Composition Design

Composition goes beyond the mechanics of type embedding and is more than just a paradigm. It is the key for maintaining stability in your software by having the ability to adapt to the data and transformation changes that are coming.

## Notes

* This is much more than the mechanics of type embedding.
* Declare types and implement workflows with composition in mind.
* Understand the problem you are trying to solve first. This means understanding the data.
* The goal is to reduce and minimize cascading changes across your software.
* Interfaces provide the highest form of composition.

## Links

http://golang.org/doc/effective_go.html#embedding  
http://www.goinggo.net/2015/09/composition-with-go.html

## Design Guidelines

* Learn about the [design guidelines](../../reading/design_guidelines.md) for composition.

## Code Review

#### Composition and Decoupling

[Struct Composition](example1/example1.go) ([Go Playground](http://play.golang.org/p/AbuUuqYhQx))  
[Decoupling With Interface](example2/example2.go) ([Go Playground](http://play.golang.org/p/oBH39i9OZv))  
[Interface Composition](example3/example3.go) ([Go Playground](http://play.golang.org/p/j55nTPKTk-))  
[Decoupling With Interface Composition](example4/example4.go) ([Go Playground](http://play.golang.org/p/zeE3PRlfFM))  

#### Conversion and Assertions

[Interface Conversions](example5/example5.go) ([Go Playground](http://play.golang.org/p/2K2svo0MR0))  
[Runtime Type Assertions](example6/example6.go) ([Go Playground](http://play.golang.org/p/tr-RGBxES-))

#### Mocking

[Package To Mock](example7/pubsub/pubsub.go) ([Go Playground](http://play.golang.org/p/3a_zYeR8M7))  
[Client](example7/example7.go) ([Go Playground](http://play.golang.org/p/guvjysMjgb))

## Exercises

### Exercise 1

Using the template, declare a set of concrete types that implement the set of predefined interface types. Then create values of these types and use them to complete a set of predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/MXFPUsqoxI)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/uXVupN6o4K))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
