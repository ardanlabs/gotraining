## Methods

Methods are functions that give data the ability to exhibit behavior.

## Notes

* Methods are functions that declare a receiver variable.
* Receivers bind a method to a type and can use value or pointer semantics.
* Value semantics mean a copy of the value is passed across program boundaries.
* Pointer semantics mean a copy of the values address is passed across program boundaries.
* Stick to a single semantic for a given type and be consistent.

## Quotes

_"Methods are valid when it is practical or reasonable for a piece of data to expose a capability." - William Kennedy_

## Links

[Methods](https://golang.org/doc/effective_go.html#methods)    
[Methods, Interfaces and Embedded Types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy    
[Escape-Analysis Flaws](https://www.ardanlabs.com/blog/2018/01/escape-analysis-flaws.html) - William Kennedy  

## Code Review

[Declare and receiver behavior](example1/example1.go) ([Go Playground](https://play.golang.org/p/-rK206XfGaV))  
[Value and Pointer semantics](example5/example5.go) ([Go Playground](https://play.golang.org/p/QmKfZAnZ6FQ))  
[Named typed methods](example2/example2.go) ([Go Playground](https://play.golang.org/p/9g1PIjyA2YQ))  
[Function/Method variables](example3/example3.go) ([Go Playground](https://play.golang.org/p/iRkiczvcHiH))  
[Function Types](example4/example4.go) ([Go Playground](https://play.golang.org/p/4TRrKs0-mTR))

## Exercises

### Exercise 1

Declare a struct that represents a baseball player. Include name, atBats and hits. Declare a method that calculates a players batting average. The formula is Hits / AtBats. Declare a slice of this type and initialize the slice with several players. Iterate over the slice displaying the players name and batting average.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/7oCUZ0IOBRK)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/smog--SovkM))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
