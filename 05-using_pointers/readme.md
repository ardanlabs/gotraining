## Nature Of Types

Think of every struct as having a nature. If the nature of the struct is something that should not be changed, like a time, a color or a coordinate, then implement the struct as a primitive data value. If the nature of the struct is something that can be changed, even if it never is in your program, it is not a primitive data value and should be implemented to be shared with a pointer. Don’t create structs that have a duality of nature.

## Notes

* The nature of the type should determine how it is shared.
* Types can implement primitive and non-primitive data.
* Don't create structs with a duality of nature.
* In general, don’t share built-in type values with a pointer.
* In general, share struct type values with a pointer unless the struct type has been implemented to behave like a primitive data value.
* In general, don’t share reference type values with a pointer unless you are implementing an unmarshal type of functionality.

## Links

http://www.goinggo.net/2014/12/using-pointers-in-go.html

## Code Review

[Primitive Types](example1/example1.go) ([Go Playground](https://play.golang.org/p/MCfjtlG9LO))

[Struct Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/zy6oUCtoSX))

[Reference Types](example3/example3.go) ([Go Playground](https://play.golang.org/p/MkfZdLcvfD))

## Exercises

### Exercise 1

Declare a struct type named Point with two fields, X and Y of type float64. Implement a factory function for this type and a method
that accept a value of this type and calculates the distance between the two points. What is the nature of this type?

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/9_MSdcdlNQ)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/5KL4HipSJ-))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
