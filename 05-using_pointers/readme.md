## Using Pointers

I like to think of types as having one of two natures. One nature a type can exhibit represents data values that should not shared. Functions and methods perform operations on these data values by accepting copies and produce new data values as their result. Values based on the built-in types exhibit this nature. User defined types can also exhibit this nature, such as types that represent values like time or coordinates. I consider these types and their data values to be primitive in nature.

The other nature a type can exhibit represents a data value that should be shared. This could be a data value that represents a document from a database. Once a data value of this type is created, passing copies to a function or method is not intuitive. Sometimes a data value of a given type is not safe to be copied and must be shared. This could be a data value that abstracts an operating system resource like a file. Passing copies of this type of data   value to a function or method can have serious consequences. I consider these types and their data values to be complex in nature.

Understanding the nature of a type will help you determine not only how to implement your methods, but also how to implement functions that create and work with values of different types.

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
