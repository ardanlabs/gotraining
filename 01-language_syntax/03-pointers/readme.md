## Pointers - Language Syntax

Pointers provide a way to share data across function boundaries. Having the ability to share and reference data with a pointer provides flexbility. It also helps our programs minimize the amount of memory they need and can add some extra performance.

## Notes

* Use pointers to share data.
* Values in Go are always pass by value.
* "Value of", what's in the box. "Address of" ( **&** ), where is the box.
* The (*) operator declares a pointer variable and the "Value that the pointer points to".

## Links

https://golang.org/doc/effective_go.html#pointers_vs_values

http://www.goinggo.net/2013/07/understanding-pointers-and-memory.html

http://www.goinggo.net/2014/12/using-pointers-in-go.html

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw)

## Code Review

[Pass by value](example1/example1.go) ([Go Playground](https://play.golang.org/p/nNnsK6hWdP))

[Sharing data I](example2/example2.go) ([Go Playground](https://play.golang.org/p/FWmGnVUDoA))

[Sharing data II](example3/example3.go) ([Go Playground](http://play.golang.org/p/VYqb11RiWr))

## Advanced Code Review

[Stack vs Heap](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/931Cw6uzsn))

## Exercises

### Exercise 1

**Part A** Declare and initalize a variable of type int with the value of 20. Display the _address of_ and _value of_ the variable.

**Part B** Declare and initialize a pointer variable of type int that points to the last variable you just created. Display the _address of_ , _value of_ and the _value that the pointer points to_.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/asM7GXfJNk)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/pUtHSTN0Ef))

### Exercise 2

Declare a struct type and create a value of this type. Declare a function that can change the value of some field in this struct type. Display the value before and after the call to your function.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/b6-FNFOToO)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/oEtveMoO1s))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).