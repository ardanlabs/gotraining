## Interface and Composition Design

Composition goes beyond the mechanics of type embedding and is more than just a paradigm. It is the key for maintaining stability in your software by having the ability to adapt to the data and transformation changes that are coming.

## Notes

* This is much more than the mechanics of type embedding.
* Declare types and implement workflows with composition in mind.
* Understand the problem you are trying to solve first. This means understanding the data.
* The goal is to reduce and minimize cascading changes across your software.
* Interfaces provide the highest form of composition.
* Don't group types by a common DNA but by a common behavior.
* Everyone can work together when we focus when we do and not what we are.

## Links

http://golang.org/doc/effective_go.html#embedding  
https://www.goinggo.net/2016/10/reducing-type-hierarchies.html  
http://www.goinggo.net/2015/09/composition-with-go.html

## Design Guidelines

* Learn about the [design guidelines](../../../reading/design_guidelines.md) for composition.

## Code Review

#### Grouping Types

[Grouping By State](grouping/example1/example1.go) ([Go Playground](https://play.golang.org/p/hKEhvhetlu))  
[Grouping By Behavior](grouping/example2/example2.go) ([Go Playground](https://play.golang.org/p/8ZsLe4nrAj))  

#### Decoupling

[Struct Composition](decoupling/example1/example1.go) ([Go Playground](https://play.golang.org/p/axLYwteYkK))  
[Decoupling With Interface](decoupling/example2/example2.go) ([Go Playground](https://play.golang.org/p/EnzMrT7Fdo))  
[Interface Composition](decoupling/example3/example3.go) ([Go Playground](https://play.golang.org/p/ES4BOnDX6O))  
[Decoupling With Interface Composition](example4/example4.go) ([Go Playground](https://play.golang.org/p/ufFSFxCdEs))  

#### Conversion and Assertions

[Interface Conversions](assertions/example1/example1.go) ([Go Playground](https://play.golang.org/p/2K2svo0MR0))  
[Runtime Type Assertions](assertions/example2/example2.go) ([Go Playground](https://play.golang.org/p/tr-RGBxES-))

#### Mocking

[Package To Mock](mocking/example1/pubsub/pubsub.go) ([Go Playground](https://play.golang.org/p/3a_zYeR8M7))  
[Client](mocking/example1/example1.go) ([Go Playground](https://play.golang.org/p/guvjysMjgb))

## Exercises

### Exercise 1

Using the template, declare a set of concrete types that implement the set of predefined interface types. Then create values of these types and use them to complete a set of predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/x6sO5GKkrs)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/XJeRRunNE2))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
