## Interface and Composition Design

Composition goes beyond the mechanics of type embedding and is more than just a paradigm. It is the key for maintaining stability in your software by having the ability to adapt to the data and transformation changes that are coming.

## Notes

* This is much more than the mechanics of type embedding.
* Declare types and implement workflows with composition in mind.
* Understand the problem you are trying to solve first. This means understanding the data.
* The goal is to reduce and minimize cascading changes across your software.
* Interfaces provide the highest form of composition.
* Don't group types by a common DNA but by a common behavior.
* Everyone can work together when we focus on what we do and not what we are.

## Quotes

_"A good API is not just easy to use but also hard to misuse." - JBD_

_"You can always embed, but you cannot decompose big interfaces once they are out there. Keep interfaces small." - JBD_

_"Don't design with interfaces, discover them." - Rob Pike_  

_"Duplication is far cheaper than the wrong abstraction. - Sandi Metz_

## Design Guidelines

* Learn about the [design guidelines](../../#interface-and-composition-design) for composition.

## Links

[Repeat yourself, do more than one thing, and rewrite everything](https://programmingisterrible.com/post/176657481103/repeat-yourself-do-more-than-one-thing-and) - tef  
[Embedding](https://golang.org/doc/effective_go.html#embedding)   
[Methods, Interfaces and Embedding](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy  
[Composition In Go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) - William Kennedy  
[Reducing Type Hierarchies](https://www.ardanlabs.com/blog/2016/10/reducing-type-hierarchies.html) - William Kennedy  
[Avoid Interface Pollution](https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html) - William Kennedy 

## Code Review

#### Grouping Types

[Grouping By State](grouping/example1/example1.go) ([Go Playground](https://play.golang.org/p/Dh_cCEz3o0N))  
[Grouping By Behavior](grouping/example2/example2.go) ([Go Playground](https://play.golang.org/p/wRpHBoPu79K))  

#### Decoupling

[Struct Composition](decoupling/example1/example1.go) ([Go Playground](https://play.golang.org/p/7nsTaKtlMWc))  
[Decoupling With Interface](decoupling/example2/example2.go) ([Go Playground](https://play.golang.org/p/HjP8V6ScpKi))  
[Interface Composition](decoupling/example3/example3.go) ([Go Playground](https://play.golang.org/p/EDbjyjjpxAi))  
[Decoupling With Interface Composition](decoupling/example4/example4.go) ([Go Playground](https://play.golang.org/p/zeO2cJLP46B))  
[Remove Interface Pollution](decoupling/example5/example5.go) ([Go Playground](https://play.golang.org/p/Kg4JKGwJGGy))  
[More Precise API](decoupling/example6/example6.go) ([Go Playground](https://play.golang.org/p/cdvbrsgclGX))

#### Conversion and Assertions

[Interface Conversions](assertions/example1/example1.go) ([Go Playground](https://play.golang.org/p/sNP3bMR1kc-))  
[Runtime Type Assertions](assertions/example2/example2.go) ([Go Playground](https://play.golang.org/p/PtdQOc9xZ7S))  
[Behavior Changes](assertions/example3/example3.go) ([Go Playground](https://play.golang.org/p/AYhH8yXDcuy))  

#### Interface Pollution

[Create Interface Pollution](pollution/example1/example1.go) ([Go Playground](https://play.golang.org/p/DCqTbY14loz))  
[Remove Interface Pollution](pollution/example2/example2.go) ([Go Playground](https://play.golang.org/p/K3w2eX7V1j2))

#### Mocking

[Package To Mock](mocking/example1/pubsub/pubsub.go) ([Go Playground](https://play.golang.org/p/299EFra4b4z))  
[Client](mocking/example1/example1.go) ([Go Playground](https://play.golang.org/p/-_laMS2yxZB))

## Exercises

### Exercise 1

Using the template, declare a set of concrete types that implement the set of predefined interface types. Then create values of these types and use them to complete a set of predefined tasks.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/6Kp_E3Wim0G)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/hwUADsRfnax))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
