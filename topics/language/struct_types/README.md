## Struct Types

Struct types are a way of creating complex types that group fields of data together. They are a great way of organizing and sharing the different aspects of the data your program consumes.

A computer architecture’s potential performance is determined predominantly by its word length (the number of bits that can be processed per access) and, more importantly, memory size, or the number of words that it can access. 

## Notes

* We can use the struct literal form to initialize a value from a struct type.
* The dot (.) operator allows us to access individual field values.
* We can create anonymous structs.

## Links

http://www.goinggo.net/2013/07/understanding-type-in-go.html  
http://www.goinggo.net/2013/07/object-oriented-programming-in-go.html  
http://dave.cheney.net/2015/10/09/padding-is-hard  
http://www.geeksforgeeks.org/structure-member-alignment-padding-and-data-packing  
http://www.catb.org/esr/structure-packing

## Code Review

[Declare, create and initialize struct types](example1/example1.go) ([Go Playground](https://play.golang.org/p/TEmOrIxl_P))  
[Anonymous struct types](example2/example2.go) ([Go Playground](https://play.golang.org/p/x-Dpp9Ts_U))  
[Named vs Unnamed types](example3/example3.go) ([Go Playground](https://play.golang.org/p/QREkSIDAuW))

## Advanced Code Review

[Struct type alignments](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/MQqKUYXoUK))

## Exercises

### Exercise 1

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/PvQKHgf9jZ)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/8CtSrnTp-1))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
