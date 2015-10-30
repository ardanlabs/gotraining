## Using Pointers

When creating a new type, try to answer this question before declaring methods. Does adding or removing something from a value of this type need to create a new value or mutate the existing one. If the answer is create a new value, then use value receivers, else use pointer receivers; and be consistent. This also applies to how values of this type should be passed to other parts of your program. Either, always use a value or a pointer, don’t mix it up. There are few exceptions to the rule.

## Notes

* The nature of the type should determine how it is passed.
* Types can implement primitive and non-primitive data qualities.
* Don't declare structs with a duality of nature.
* In general, don’t pass built-in type values with a pointer.
* In general, don’t pass reference type values with a pointer unless you are implementing unmarshal type of functionality.
* In general, pass struct type values with a pointer unless the struct type has been implemented to behave like a primitive data value.

## Links

http://www.goinggo.net/2014/12/using-pointers-in-go.html

http://play.golang.org/p/ki991PuHhk

## Code Review

[Primitive Types](example1/example1.go) ([Go Playground](https://play.golang.org/p/H5HRoElN6q))

[Reference Types](example3/example3.go) ([Go Playground](https://play.golang.org/p/E-Bb5cRuyz))

[Struct Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/xD6PCx--GG))

## Exercises

### Exercise 1

Declare a struct type named Point with two fields, X and Y of type int. Implement a factory function for this type and a method
that accepts this type and calculates the distance between the two points. What is the nature of this type?

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/Qe3dhDTwzX)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/1tcEsqNG6a))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
