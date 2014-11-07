## Interfaces - Methods, Interfaces and Embedding

Interfaces provide a way to declare types that define behavior. Then struct and named type can declare methods and implement this behavior. When a struct or named type implements the behavior declare by an interface, it satisfies the interface and can be assigned as values of the interface type. This leads to providing polymorphic behavior in our programs.

## Notes

* The method set for a value, only includes methods implemented with a value reciever.
* The method set for a pointer, includes methods implemented with both pointer and value recievers.
* Interface values with an underlying value, can only call methods implemented with a value receiver.
* Interface values with an underlying pointer, can call methods implemented with both pointer and value receivers.
* The rules for method calls with concrete typed values do not apply.
* Interfaces are reference types, don't share with a pointer.
* This is how we create polymorphic behavior in go.

## Links

https://golang.org/doc/effective_go.html#interfaces

http://blog.golang.org/laws-of-reflection

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## Code Review

[Declare, implement and method call restrictions I](example1/example1.go) ([Go Playground](http://play.golang.org/p/h5Q9dQgnzS))

[Declare, implement and method call restrictions II](example2/example2.go) ([Go Playground](http://play.golang.org/p/byYKqtmHFU))

## Exercises

### Exercise 1

**Part A** Declare an interface named Speaker with a method named SayHello. Declare a struct named English that represents a person who speaks english and declare a struct named Chinese for someone who speaks chinese. Implement the Speaker interface for each struct using a pointer receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type Speaker and assign the _address of_ a value of type English and call the method. Do it again for a value of type Chinese.

**Part B** From exercise 1, add a new function named SayHello that accepts a value of type Speaker. Implement that function to call the SayHello method on the interface value. Then create new values of each type and use the function.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/pbcD5WmTX9))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).