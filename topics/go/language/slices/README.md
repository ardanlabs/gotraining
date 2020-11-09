## Slices - Arrays, Slices and Maps

Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

## Notes

* Slices are like dynamic arrays with special and built-in functionality.
* There is a difference between a slices length and capacity and they each service a purpose.
* Slices allow for multiple "views" of the same underlying array.
* Slices can grow through the use of the built-in function append.

## Links

[Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) - Andrew Gerrand    
[Strings, bytes, runes and characters in Go](https://blog.golang.org/strings) - Rob Pike    
[Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices) - Rob Pike    
[Understanding Slices in Go Programming](https://www.ardanlabs.com/blog/2013/08/understanding-slices-in-go-programming.html) - William Kennedy    
[Collections Of Unknown Length in Go](https://www.ardanlabs.com/blog/2013/08/collections-of-unknown-length-in-go.html) - William Kennedy    
[Iterating Over Slices In Go](https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html) - William Kennedy    
[Slices of Slices of Slices in Go](https://www.ardanlabs.com/blog/2013/09/slices-of-slices-of-slices-in-go.html) - William Kennedy    
[Three-Index Slices in Go 1.2](https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html) - William Kennedy    
[SliceTricks](https://github.com/golang/go/wiki/SliceTricks)    

## Code Review

[Declare and Length](example1/example1.go) ([Go Playground](https://play.golang.org/p/ydOJ1GHgR_Y))  
[Reference Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/WqDnss06_9E))  
[Appending slices](example4/example4.go) ([Go Playground](https://play.golang.org/p/E-NTGM6daAA))  
[Taking slices of slices](example3/example3.go) ([Go Playground](https://play.golang.org/p/rUP9grCot8J))  
[Slices and References](example5/example5.go) ([Go Playground](https://play.golang.org/p/D88zzGYanvX))  
[Strings and slices](example6/example6.go) ([Go Playground](https://play.golang.org/p/1RntHk6UPA5))  
[Variadic functions](example7/example7.go) ([Go Playground](https://play.golang.org/p/rUjWVBMmxgP))  
[Range mechanics](example8/example8.go) ([Go Playground](https://play.golang.org/p/d1wToBg6oUu))  
[Efficient Traversals](example9/example9.go) ([Go Playground](https://play.golang.org/p/xPL8U_bD4kD))  

## Advanced Code Review

[Three index slicing](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/2CM_LPBnfIR))

## Exercises

### Exercise 1

**Part A** Declare a nil slice of integers. Create a loop that appends 10 values to the slice. Iterate over the slice and display each value.

**Part B** Declare a slice of five strings and initialize the slice with string literal values. Display all the elements. Take a slice of index one and two and display the index position and value of each element in the new slice.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/7GfB3NOwu_c)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/0xv7GTHHIR_K))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
