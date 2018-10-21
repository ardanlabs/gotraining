## Exporting

Packages contain the basic unit of compiled code. They define a scope for the identifiers that are declared within them. Exporting is not the same as public and private semantics in other languages. But exporting is how we provide encapsulation in Go.

## Notes

* Code in go is compiled into packages and then linked together.
* Identifiers are exported (or remain unexported) based on letter-case.
* We import packages to access exported identifiers.
* Any package can use a value of an unexported type, but this is annoying to use.

## Links

[Exported/Unexported Identifiers In Go](https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html) - William Kennedy    

## Code Review

[Declare and access exported identifiers - Pkg](example1/counters/counters.go) ([Go Playground](https://play.golang.org/p/8Xzq-m9ez-I))  
[Declare and access exported identifiers - Main](example1/example1.go) ([Go Playground](https://play.golang.org/p/KrpX0CyIyYO))  

[Declare unexported identifiers and restrictions - Pkg](example2/counters/counters.go) ([Go Playground](https://play.golang.org/p/9u1IQexx5gk))  
[Declare unexported identifiers and restrictions - Main](example2/example2.go) ([Go Playground](https://play.golang.org/p/A5FpmRpuOWJ))  

[Access values of unexported identifiers - Pkg](example3/counters/counters.go) ([Go Playground](https://play.golang.org/p/NroO30yoNvh))  
[Access values of unexported identifiers - Main](example3/example3.go) ([Go Playground](https://play.golang.org/p/e5fg0uOEkkn))  

[Unexported struct type fields - Pkg](example4/users/users.go) ([Go Playground](https://play.golang.org/p/KQ6x5z7E1pN))  
[Unexported struct type fields - Main](example4/example4.go) ([Go Playground](https://play.golang.org/p/6MznWaiGwr-))  

[Unexported embedded types - Pkg](example5/users/users.go) ([Go Playground](https://play.golang.org/p/br-2rVc1VF1))  
[Unexported embedded types - Main](example5/example5.go) ([Go Playground](https://play.golang.org/p/p9pQo5gCB42))  

## Exercises

### Exercise 1
**Part A** Create a package named toy with a single exported struct type named Toy. Add the exported fields Name and Weight. Then add two unexported fields named onHand and sold. Declare a factory function called New to create values of type toy and accept parameters for the exported fields. Then declare methods that return and update values for the unexported fields.

**Part B** Create a program that imports the toy package. Use the New function to create a value of type toy. Then use the methods to set the counts and display the field values of that toy value.

[Template](exercises/template1) |
[Answer](exercises/exercise1)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
