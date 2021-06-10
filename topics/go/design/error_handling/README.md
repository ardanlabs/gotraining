## Error Handling Design

Error handling is critical for making your programs reliable, trustworthy and respectful to those who depend on them. A proper error value is both specific and informative. It must allow the caller to make an informed decision about the error that has occurred. There are several ways in Go to create error values. This depends on the amount of context that needs to be provided.

## Notes

* Use the default error value for static and simple formatted messages.
* Create and return error variables to help the caller identify specific errors.
* Create custom error types when the context of the error is more complex.
* Error Values in Go aren't special, they are just values like any other, and so you have the entire language at your disposal.

## Quotes

_Systems cannot be developed assuming that human beings will be able to write millions of lines of code without making mistakes, and debugging alone is not an efficient way to develop reliable systems. - Al Aho (inventor of AWK)_

## Links

[Error handling and Go](https://blog.golang.org/error-handling-and-go)    
[Error Handling In Go, Part I](https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html) - William Kennedy    
[Error Handling In Go, Part II](https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html) - William Kennedy    
[Design Philosophy On Logging](https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html) - William Kennedy    
[Bugs are a failure of prediction](https://clipperhouse.com/bugs-are-a-failure-of-prediction/) - Matt Sherman    
[Inspecting errors](https://dave.cheney.net/2014/12/24/inspecting-errors) - Dave Cheney    
[Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) - Dave Cheney    
[Stack traces and the errors package](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package) - Dave Cheney    
[Errors are handled in return values](https://plus.google.com/+RussCox-rsc/posts/iqAiKAwP6Ce) - Russ Cox    
[Error handling in Upspin](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html) - Rob Pike    
[Go 1.13: xerrors](https://crawshaw.io/blog/xerrors) - David Crawshaw  
[Why Go's Error Handling is Awesome](https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html) - Raul Jordan  

## Code Review

[Default Error Values](example1/example1.go) ([Go Playground](https://play.golang.org/p/beGEdO2QE4g))  
[Error Variables](example2/example2.go) ([Go Playground](https://play.golang.org/p/JQUJbS20MrE))  
[Type As Context](example3/example3.go) ([Go Playground](https://play.golang.org/p/BmiblC2Q7MC))  
[Behavior As Context](example4/example4.go) ([Go Playground](https://play.golang.org/p/sNRSXKtcJKM))  
[Find The Bug](example5/example5.go) ([Go Playground](https://play.golang.org/p/CBL-ADH-nSv)) | 
[The Reason](example5/reason/reason.go) ([Go Playground](https://play.golang.org/p/-f4PPcBGkDU))  
[Wrapping Errors With pkg/errors](example6/example6.go) ([Go Playground](https://play.golang.org/p/Zt1Z5k4HbDG))  
[Wrapping Errors With stdlib](example7/example7.go) ([Go Playground](https://play.golang.org/p/f5bw9G7OLog))  

## Exercises

### Exercise 1
Create two error variables, one called ErrInvalidValue and the other called ErrAmountTooLarge. Provide the static message for each variable. Then write a function called checkAmount that accepts a float64 type value and returns an error value. Check the value for zero and if it is, return the ErrInvalidValue. Check the value for greater than $1,000 and if it is, return the ErrAmountTooLarge. Write a main function to call the checkAmount function and check the return error value. Display a proper message to the screen.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/uOAUy1AoP6t)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/OEq2bZ-fcJZ))

### Exercise 2
Create a custom error type called appError that contains three fields, err error, message string and code int. Implement the error interface providing your own message using these three fields. Implement a second method named temporary that returns false when the value of the code field is 9. Write a function called checkFlag that accepts a bool value. If the value is false, return a pointer of your custom error type initialized as you like. If the value is true, return a default error. Write a main function to call the checkFlag function and check the error using the temporary interface.

[Template](exercises/template2/template2.go) ([Go Playground](https://play.golang.org/p/SqrcJVqwT1X)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](https://play.golang.org/p/EzdPD58tQ4D))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
