## Type Conversions

Go is a strict type safe language so typed values must be explicitly converted from one type to the other. Untyped constants on the other hand can be implicitly converted by the compiler. These values exist in a kind system and have more flexibility. In these examples we will learn about named types and how the time package uses them, in conjuction with constants, to provide both a flexible and type safe API.

## Notes

* Declare a type based on another type including built-in and user defined types.
* Learn about explicit and implicit conversions.
* See the power of constants and their use in the standard library.
* Use typed constants are part of the API for your package functions and methods.

## Code Review

[Declare, create and initialize named types](example1/example1.go) ([Go Playground](http://play.golang.org/p/jcNHcgFz6N))  
[Named types in the standard library](example2/example2.go) ([Go Playground](http://play.golang.org/p/bUtxOEgcOn))  
[Conversions I](example3/example3.go) ([Go Playground](http://play.golang.org/p/NXJ8i2Gkhc))  
[Conversions II](example4/example4.go) ([Go Playground](http://play.golang.org/p/cc3XhBsKVW))

## Exercises

### Exercise 1

**Part A** Declare a named type called counter with a base type of int. Declare and initialize a variable of this named type to its zero value. Display the value of this variable and the variables type.

**Part B** Declare a new variable of the named type assign it the value of 10. Display the value.

**Part C** Declare a variable of the same base type as your named typed. Attempt to assign the value of your named type variable to your new base type variable. Does the compiler allow the assignment?

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/stz8qh6YeR)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/aL6INg8hAh))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
