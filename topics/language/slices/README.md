## Slices - Arrays, Slices and Maps

Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

## Notes

* Slices are like dynamic arrays with special and built-in functionality.
* There is a difference between a slices length and capacity and they each service a purpose.
* Slices allow for multiple "views" of the same underlying array.
* Slices can grow through the use of the built-in function append.

## Links

http://blog.golang.org/go-slices-usage-and-internals  
https://blog.golang.org/strings  
http://blog.golang.org/slices  
http://www.goinggo.net/2013/08/understanding-slices-in-go-programming.html  
http://www.goinggo.net/2013/08/collections-of-unknown-length-in-go.html  
http://www.goinggo.net/2013/09/iterating-over-slices-in-go.html  
http://www.goinggo.net/2013/09/slices-of-slices-of-slices-in-go.html  
http://www.goinggo.net/2013/12/three-index-slices-in-go-12.html

## Code Review

[Declare and Length](example1/example1.go) ([Go Playground](https://play.golang.org/p/lDKravTEqF))  
[Reference Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/gVWb35XjwM))  
[Appending slices](example4/example4.go) ([Go Playground](https://play.golang.org/p/x4hGDV1phU))  
[Taking slices of slices](example3/example3.go) ([Go Playground](https://play.golang.org/p/Okc2EZG5_M))  
[Slices and References](example5/example5.go) ([Go Playground](https://play.golang.org/p/XGrtsrxmNb))  
[Strings and slices](example6/example6.go) ([Go Playground](https://play.golang.org/p/x0Q5ByzxGS))  
[Variadic functions](example7/example7.go) ([Go Playground](https://play.golang.org/p/aTGRT1rhoO))

## Advanced Code Review

[Three index slicing](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/QZQIdaTgtG))

## Exercises

### Exercise 1

**Part A** Declare a nil slice of integers. Create a loop that appends 10 values to the slice. Iterate over the slice and display each value.

**Part B** Declare a slice of five strings and initialize the slice with string literal values. Display all the elements. Take a slice of index one and two and display the index position and value of each element in the new slice.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/sE06PRtw7h)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/3WKISOXA-L))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
