## Constants - Language Syntax

Constants are a way to create a named identifier who's value can never change. They also provide an incredible amount of flexibility to the language. The way constants are implemented in Go is very unique.

## Notes

* Constants are not variables.
* They existing only at compilation.
* Untyped constants can be implictly converted where typed constants and variables can't.
* Think of untyped constants as having a Kind, not a Type.

## Links

https://golang.org/ref/spec#Constants

http://blog.golang.org/constants

http://www.goinggo.net/2014/04/introduction-to-numeric-constants-in-go.html

## Code Review

[Declare and inialize constants](example1/example1.go) ([Go Playground](http://play.golang.org/p/OLuzwK1oHT))

[Parallel type system (Kind)](example2/example2.go) ([Go Playground](http://play.golang.org/p/ExxRWe6jUz))

## Exercises

### Exercise 1

**Part A:** Declare an untyped and typed constant and display their values.

**Part B:** Multiply two literal constants into a typed variable and display the value.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/44wgDZ-U2t))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).