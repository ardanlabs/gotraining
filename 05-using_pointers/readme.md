## Using Pointers

I like to think of types as having one of two natures. One nature is a type that represents a data value that should not be shared. Data values that are created using a built-in or reference type exhibit this primitive nature. These data values should always be passed using copies of the original. The other nature is a type that should be shared. Data values that are created using struct types exhibit this nature in most cases. These data values should always be passed by sharing them with a pointer.

However, struct types can also exhibit a primitive nature like the built-in and reference types do. Struct types that represent time or coordinate data values are a good example of this. Understanding the nature of a type will help you determine how best to pass your data values between methods and functions.

## Notes

* The nature of the type should determine how it is passed.
* Types can implement primitive and non-primitive data.
* Don't create structs with a duality of nature.
* In general, don’t pass built-in type values with a pointer.
* In general, pass struct type values with a pointer unless the struct type has been implemented to behave like a primitive data value.
* In general, don’t pass reference type values with a pointer unless you are implementing an unmarshal type of functionality.

## Links

http://www.goinggo.net/2014/12/using-pointers-in-go.html

http://play.golang.org/p/ki991PuHhk

## Code Review

[Primitive Types](example1/example1.go) ([Go Playground](https://play.golang.org/p/H5HRoElN6q))

[Struct Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/xD6PCx--GG))

[Reference Types](example3/example3.go) ([Go Playground](https://play.golang.org/p/E-Bb5cRuyz))

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
