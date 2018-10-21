## Constants

Constants are a way to create a named identifier whose value can never change. They also provide an incredible amount of flexibility to the language. The way constants are implemented in Go is very unique.

## Notes

* Constants are not variables.
* They exist only at compilation.
* Untyped constants can be implicitly converted where typed constants and variables can't.
* Think of untyped constants as having a Kind, not a Type.
* Learn about explicit and implicit conversions.
* See the power of constants and their use in the standard library.

## Links

[Constants specification](https://golang.org/ref/spec#Constants)    
[Constants](https://blog.golang.org/constants) - Rob Pike    
[Introduction To Numeric Constants In Go](https://www.ardanlabs.com/blog/2014/04/introduction-to-numeric-constants-in-go.html) - William Kennedy    

## Code Review

[Declare and initialize constants](example1/example1.go) ([Go Playground](https://play.golang.org/p/z251qax3MYa))  
[Parallel type system (Kind)](example2/example2.go) ([Go Playground](https://play.golang.org/p/8a_tp97RHAf))  
[iota](example3/example3.go) ([Go Playground](https://play.golang.org/p/SLAYYNFIdUA))  
[Implicit conversion](example4/example4.go) ([Go Playground](https://play.golang.org/p/aB4NGcnZlw2))  

## Exercises

### Exercise 1

**Part A:** Declare an untyped and typed constant and display their values.

**Part B:** Divide two literal constants into a typed variable and display the value.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/4Gs3Ls_5_pi)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/Znc6RAvrF_c))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
