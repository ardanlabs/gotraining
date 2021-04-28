## Interfaces

Interfaces provide a way to declare types that define only behavior. This behavior can be implemented by concrete types, such as struct or named types, via methods. When a concrete type implements the set of methods for an interface, values of the concrete type can be assigned to variables of the interface type. Then method calls against the interface value actually call into the equivalent method of the concrete value. Since any concrete type can implement any interface, method calls against an interface value are polymorphic in nature.

## Notes

* The method set for a value, only includes methods implemented with a value receiver.
* The method set for a pointer, includes methods implemented with both pointer and value receivers.
* Methods declared with a pointer receiver, only implement the interface with pointer values.
* Methods declared with a value receiver, implement the interface with both a value and pointer receiver.
* The rules of method sets apply to interface types.
* Interfaces are reference types, don't share with a pointer.
* This is how we create polymorphic behavior in go.

## Quotes

_"Polymorphism means that you write a certain program and it behaves differently depending on the data that it operates on." - Tom Kurtz (inventor of BASIC)_

_"The empty interface says nothing." - Rob Pike_

_"Design is the art of arranging code to work today, and be changeable forever." - Sandi Metz_

_"A proper abstraction decouples the code so that every change doesn’t echo throughout the entire code base."_ - Ronna Steinburg

## Links

[Interfaces](https://golang.org/doc/effective_go.html#interfaces)    
[The Laws of Reflection](https://blog.golang.org/laws-of-reflection) - Rob Pike    
[Methods, Interfaces and Embedded Types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy    
[Interface Pollution](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275) - JBD    
[Abstraction Considered Harmful](https://bravenewgeek.com/abstraction-considered-harmful/) - Tyler Treat    
[Interface Values Are Valueless](https://www.ardanlabs.com/blog/2018/03/interface-values-are-valueless.html) - William Kennedy    
[Interface Semantics](https://www.ardanlabs.com/blog/2017/07/interface-semantics.html) - William Kennedy    
[Hyrum's Law](https://www.hyrumslaw.com/) - Hyrum  

## Code Review

[Repetitive Code That Needs Polymorphism](example0/example0.go) ([Go Playground](https://play.golang.org/p/Txsuzcpdran))  
[Polymorphism](example1/example1.go) ([Go Playground](https://play.golang.org/p/J7OWzPzV40w))  
[Method Sets](example2/example2.go) ([Go Playground](https://play.golang.org/p/N50ocjUekf3))  
[Address Of Value](example3/example3.go) ([Go Playground](https://play.golang.org/p/w981JSUcVZ2))  
[Storage By Value](example4/example4.go) ([Go Playground](https://play.golang.org/p/6KE5b1EwovK))  
[Type Assertions](example5/example5.go) ([Go Playground](https://play.golang.org/p/f47JMTj2eId))  
[Conditional Type Assertions](example6/example6.go) ([Go Playground](https://play.golang.org/p/9fYc5RyyvVG))  
[The Empty Interface and Type Switches](example7/example7.go) ([Go Playground](https://play.golang.org/p/iyDfKCIQ4S9))  

## Advanced Code Review

[Storing Values](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/yDK5lUiPPHW))

## Exercises

### Exercise 1

**Part A** Declare an interface named speaker with a method named speak. Declare a struct named english that represents a person who speaks english and declare a struct named chinese for someone who speaks chinese. Implement the speaker interface for each struct using a value receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type speaker and assign the address of a value of type english and call the method. Do it again for a value of type chinese.

**Part B** Add a new function named sayHello that accepts a value of type speaker. Implement that function to call the speak method on the interface value. Then create new values of each type and use the function.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/hfC2-yPI9y6)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/mN8Fitr8Wts))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
