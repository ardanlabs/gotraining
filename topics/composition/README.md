## Interface and Composition Design

Composition goes beyond the mechanics of type embedding and is more than just a paradigm. It is the key for maintaining stability in your software by having the ability to adapt to the data and transformation changes that are coming.

## Notes

* This is much more than the mechanics of type embedding.
* Declare types and implement workflows with composition in mind.
* Understand the problem you are trying to solve first. This means understanding the data.
* The goal is to reduce and minimize cascading changes across your software.
* Interfaces provide the highest form composition.

## Links

http://golang.org/doc/effective_go.html#embedding

http://www.goinggo.net/2015/09/composition-with-go.html

## Design Guidelines

* Learn about the [design guidelines](../../reading/design_guidelines.md) for composition.

## Code Review

[Struct Composition](example1/example1.go) ([Go Playground](https://play.golang.org/p/wipPTC9se1))

[Decoupling With Interface](example2/example2.go) ([Go Playground](https://play.golang.org/p/Kh8JCDxdjY))

[Interface Composition](example3/example3.go) ([Go Playground](https://play.golang.org/p/wUtZ7gxLIL))

[Decoupling With Interface Composition](example4/example4.go) ([Go Playground](https://play.golang.org/p/uB4c33sbfj))

[Interface Conversions](example5/example5.go) ([Go Playground](http://play.golang.org/p/W8_QflbEFz))

[Runtime Type Assertions](example6/example6.go) ([Go Playground](http://play.golang.org/p/2kfVP_SGA4))

## Exercises

### Exercise 1

**Part A** Follow the guided comments to:

**Part B** Declare a sysadmin type that implements the administrator interface.

**Part C** Declare a programmer type that implements the developer interface.

**Part D** Declare a company type that embeds both an administrator and a developer.

**Part E** Create a sysadmin, programmers, and a company which are available for hire, and use them to complete some predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/b8ww3jd2Xs)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/UvFEZQHDu0))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).