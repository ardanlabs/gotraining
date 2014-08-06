## Slices
Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

### Code Review

[Declare and Length](example1/example1.go) ([Go Playground](http://play.golang.org/p/4r90uFQwJn))

[Reference Types](example2/example2.go) ([Go Playground](http://play.golang.org/p/DB8hwJ0hw9)

[Taking slices of slices](example3/example3.go) ([Go Playground](http://play.golang.org/p/PyZthd9EFl))

[Appending slices](example4/example4.go) ([Go Playground](http://play.golang.org/p/UzmwiMWDwd))

[Iterating over slices](example5/example5.go) ([Go Playground](http://play.golang.org/p/HV5t0VrRie))

[Three index slicing](example6/example6.go) ([Go Playground](http://play.golang.org/p/v3ZHknDvSx))

(Advanced) [Practical use of slices](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/-qQgO7NbLm))

### Exercise 1
Declare a nil slice of integers. Create a loop that increments a counter variable by 10 five times and appends these values to the slice. Iterate over the slice and display each value.

### Exercise 2
Declare a slice of five strings and initialize the slice with string literal values. Take a slice of the first, second and third elements and display the index position and value of each element for both slices.

### Exercise 3
Declare a slice of five strings and initialize this slice with string literal values. Take a slice of the second element with a capacity of one. Display the length and capacity of the slice. Iterate over the slice and display the address and value of each element. Append a new value to the slice and display everything again.

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)