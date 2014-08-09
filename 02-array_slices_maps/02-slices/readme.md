## Slices - Arrays, Slices and Maps

Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

## Notes

* Slices are like dynamic arrays with special and built-in functionality.
* There is a difference between a slices length and capacity and they each service a purpose.
* Slices allow for multiple "views" of the same underlying array.

## Code Review

[Declare and Length](example1/example1.go) ([Go Playground](http://play.golang.org/p/4r90uFQwJn))

[Reference Types](example2/example2.go) ([Go Playground](http://play.golang.org/p/DB8hwJ0hw9)

[Taking slices of slices](example3/example3.go) ([Go Playground](http://play.golang.org/p/vlRlYsfLwb))

[Appending slices](example4/example4.go) ([Go Playground](http://play.golang.org/p/UzmwiMWDwd))

[Iterating over slices](example5/example5.go) ([Go Playground](http://play.golang.org/p/HV5t0VrRie))

## Advanced Code Review

[Practical use of slices](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/-qQgO7NbLm))

[Three index slicing](advanced/example2/example2.go) ([Go Playground](http://play.golang.org/p/v3ZHknDvSx))

## Exercies

### Exercise 1

**Part A** Declare a nil slice of integers. Create a loop that increments a counter variable by 10 five times and appends these values to the slice. Iterate over the slice and display each value.

**Part B** Declare a slice of five strings and initialize the slice with string literal values. Display all the elements. Take a slice of the second and third elements (index 1 and 2) and display the index position and value of each element in the new slice.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/9ir4vinceh))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)