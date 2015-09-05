## Arrays - Arrays, Slices and Maps

Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory. Arrays have some special features in Go related to how they are declared and viewed as types.

## Notes

* Arrays are fixed length data structures that can't change.
* Arrays of different sizes are considered to be of different types.
* Memory is allocated as a contiguous block.

## Code Review

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](http://play.golang.org/p/2D24t6fbW_))

[Different type arrays](example2/example2.go) ([Go Playground](http://play.golang.org/p/nYgwqqOctt))

[Contiguous memory allocations](example3/example3.go) ([Go Playground](http://play.golang.org/p/L-SmdGfUcP))

## Exercises

### Exercise 1

Declare an array of 5 strings with each element initialized to its zero value. Declare a second array of 5 strings and initialize this array with literal string values. Assign the second array to the first and display the results of the first array.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/ggjjRPzhAB)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/Pa3mrTCcpB))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).