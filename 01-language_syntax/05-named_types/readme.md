## Named Types - Language Syntax

Named types are another way to create user defined types. They provide some interesting feature/functionality that is not always available in other languages. The standard library in Go used named types very effectively. In these examples a named type is used to declare a set of typed constants to act as part of the API for the time package.

## Notes

* Declare a type based on another single type including built-in and user defined types.
* See the power of constants and their use in the standard library.
* Use typed constants are part of the API for your package functions and methods.

## Code Review

[Declare, create and initalize named types](example1/example1.go) ([Go Playground](http://play.golang.org/p/mhKlxSyuxr))

[Named types in the standard library](example2/example2.go) ([Go Playground](http://play.golang.org/p/XJ4Ia1lMWl))

[Named types and conversion I](example3/example3.go) ([Go Playground](http://play.golang.org/p/Rgoqvg8dNv))

[Named types and conversion II](example4/example4.go) ([Go Playground](http://play.golang.org/p/B75FURdQ7t))

## Exercises

### Exercise 1

**Part A** Declare a named type called counter with a base type of int. Declare and initalize a variable of this named type to its zero value. Display the value of this variable and the variables type.

**Part B** Declare a new variable of the named type assign it the value of 10. Display the value.

**Part C** Declare a variable of the same base type as your named typed. Attempt to assign the value of your named type variable to your new base type variable. Does the compiler allow the assignment?

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/_zBf2HVsAU))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)