## Functions

Functions are at the core of the language. They provide a mechanism to group and organize our code to separate and distinct pieces of functionality. They can be used to provide an API to the packages we write and are a core component to concurrency.

## Notes

* Functions can return multiple values and most return an error value.
* The error value should always be checked as part of the programming logic.
* The blank identifier can be used to ignore return values.

## Links

https://golang.org/doc/effective_go.html#functions  
http://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html  
http://www.goinggo.net/2013/06/understanding-defer-panic-and-recover.html

## Code Review

[Return multiple values](example1/example1.go) ([Go Playground](https://play.golang.org/p/rJMtATFqPi))  
[Blank identifier](example2/example2.go) ([Go Playground](https://play.golang.org/p/ziCWrNaGWO))  
[Redeclarations](example3/example3.go) ([Go Playground](https://play.golang.org/p/CofPHyVpne))  
[Anonymous Functions/Closures](example4/example4.go) ([Go Playground](https://play.golang.org/p/AhT35gu2fE))

## Advanced Code Review

[Recover panics](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/UuT3FNWd7x))

## Exercises

### Exercise 1

**Part A** Declare a struct type to maintain information about a user. Declare a function that creates value of and returns pointers of this type and an error value. Call this function from main and display the value.

**Part B** Make a second call to your function but this time ignore the value and just test the error value.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/i5wI736jpN)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/fabhfnqJ0C))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
