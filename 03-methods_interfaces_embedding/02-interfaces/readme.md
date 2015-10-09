## Interfaces - Methods, Interfaces and Embedding

Interfaces provide a way to declare types that define only behavior. This behavior can be implemented by concrete types, such as struct or named types, via methods. When a concrete type implements the set of methods for an interface, values of the concrete type can be assigned to variables of the interface type. Then method calls against the interface value actually call into the equivalent method of the concrete value. Since any concrete type can implement any interface, method calls against an interface value are polymorphic in nature.

## Notes

* The method set for a value, only includes methods implemented with a value reciever.
* The method set for a pointer, includes methods implemented with both pointer and value recievers.
* Methods declared with a pointer receiver, only implement the interface with pointer values.
* Methods declared with a value receiver, implement the interface with both a value and pointer receiver.
* The rules of method sets apply to interface types.
* Interfaces are reference types, don't share with a pointer.
* This is how we create polymorphic behavior in go.

## Links

https://golang.org/doc/effective_go.html#interfaces

http://blog.golang.org/laws-of-reflection

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275

## Code Review

[Method Receivers](example1/example1.go) ([Go Playground](http://play.golang.org/p/xvOEbidmSQ))

[Polymorphism](example2/example2.go) ([Go Playground](http://play.golang.org/p/cjZ-iIHIsK))

[Address Of Value](example3/example3.go) ([Go Playground](https://play.golang.org/p/lTMxc-oExx))

## Exercises

### Exercise 1

**Part A** Declare an interface named Speaker with a method named SayHello. Declare a struct named English that represents a person who speaks english and declare a struct named Chinese for someone who speaks chinese. Implement the Speaker interface for each struct using a value receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type Speaker and assign the _address of_ a value of type English and call the method. Do it again for a value of type Chinese.

**Part B** Add a new function named SayHello that accepts a value of type Speaker. Implement that function to call the SayHello method on the interface value. Then create new values of each type and use the function.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/oijJdRW3cD)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/MFPF58SkRs))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).