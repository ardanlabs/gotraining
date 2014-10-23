## Functions - Language Syntax

Function are at the core of the language. They provide a mechanism to group and organize our code to separate and distinct pieces of functionality. They can be used to provide an API to the packages we write and are a core compontent to concurrency.

## Notes

* Functions can return mulitple values and most return an error value.
* The error value should always be checked as part of the programming logic.
* The blank identifier can be used to ignore return values.

## Links

https://golang.org/doc/effective_go.html#functions

http://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html

http://www.goinggo.net/2013/06/understanding-defer-panic-and-recover.html

## Code Review

[Return multiple values](example1/example1.go) ([Go Playground](http://play.golang.org/p/bYY-TRjfH0))

[Blank identifier](example2/example2.go) ([Go Playground](http://play.golang.org/p/wPVvgwPlHw))

[Variadic functions](example3/example3.go) ([Go Playground](http://play.golang.org/p/cK3y_qYUgd))

## Advanced Code Review

[Trapping panics](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/eg14ClW4_y))

## Exercises

### Exercise 1

**Part A** Declare a struct type to maintain information about a user. Declare a function that creates value of and returns pointers of this type and an error value. Call this function from main and display the value.

**Part B** Make a second call to your function but this time ignore the value and just test the error value.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/hJ532ON6Xb))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).