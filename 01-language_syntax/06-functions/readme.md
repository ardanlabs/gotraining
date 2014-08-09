## Functions - Language Syntax

Function are at the core of the language. They provide a mechanism to group and organize our code to separate and distinct pieces of functionality. They can be used to provide an API to the packages we write and are a core compontent to concurrency.

## Notes

* Functions can return mulitple values are most return an error value.
* The error value should always be checked as part of the programming logic.
* The blank identifier can be used to ignore return values.

## Code Review

[Return multiple values](example1/example1.go) ([Go Playground](http://play.golang.org/p/kTKdUJolAU))

[Blank identifier](example2/example2.go) ([Go Playground](http://play.golang.org/p/dDZpl7ti1I))

[Variadic functions](example3/example3.go) ([Go Playground](http://play.golang.org/p/RoP6pNPgKl))

## Advanced Code Review

[Trapping panics](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/eg14ClW4_y))

## Exercises

### Exercise 1

**Part A** Declare a struct type to maintain information about a user. Declare a function that creates value of and returns pointers of this type and an error value. Call this function from main and display the value.

**Part B** Make a second call to your function but this time ignore the value and just test the error value.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](NEED PLAYGROUND))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)