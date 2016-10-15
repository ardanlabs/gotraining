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

## Interface Design Guidelines

* Learn about the [design guidelines](../../reading/design_guidelines.md) for interfaces.

## Links

https://golang.org/doc/effective_go.html#interfaces  
http://blog.golang.org/laws-of-reflection  
http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html  
[Interface Pollution](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275)  
[Abstraction Considered Harmful](http://bravenewgeek.com/abstraction-considered-harmful)

## Code Review

[Polymorphism](example1/example1.go) ([Go Playground](https://play.golang.org/p/hbz_OvJD_p))  
[Method Sets](example2/example2.go) ([Go Playground](https://play.golang.org/p/4R3_QVKNli))  
[Address Of Value](example3/example3.go) ([Go Playground](https://play.golang.org/p/hJtuUbNICG))  
[Behavior Changes](example4/example4.go) ([Go Playground](https://play.golang.org/p/OrFNjhTrxv))  
[Storage By Value](example5/example5.go) ([Go Playground](https://play.golang.org/p/9yHyRQUEkW))  

## Advanced Code Review

[Storing Values](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/KXvtpd9_29))

## Exercises

### Exercise 1

**Part A** Declare an interface named speaker with a method named speak. Declare a struct named english that represents a person who speaks english and declare a struct named chinese for someone who speaks chinese. Implement the speaker interface for each struct using a value receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type speaker and assign the address of a value of type english and call the method. Do it again for a value of type chinese.

**Part B** Add a new function named sayHello that accepts a value of type speaker. Implement that function to call the speak method on the interface value. Then create new values of each type and use the function.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/adkJ3OvYpr)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/06fecJbfE4))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
