## Embedding - Methods, Interfaces and Embedding

Embedding types provides the final piece of sharing and reusing state and behavior between types. Through the use of inner type promotion, an inner type's fields and methods can be directly access by references of the outer type.

## Notes

* Embedding types allow us to share state or behavior between types.
* The inner type never loses its identity.
* This is not inheritance.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* The outer type can override the inner type's behavior.

## Links

http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html

## Code Review

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/e5O_Dx5VpM))

[Embedding types](example2/example2.go) ([Go Playground](https://play.golang.org/p/UkrDXkk-Ch))

[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/BgEoThS7u9))

[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/jfOfrRMPZR))

## Exercises

### Exercise 1

**Part A** Declare a struct type named animal with two fields name and age. Declare a struct type named dog with the field bark. Embed the animal type into the dog type. Declare and initalize a value of type dog. Display the value of the variable.

**Part B** Add a method to the animal type using a pointer reciever named yelp which displays the literal string "Not Implemented". Call the method from the value of type dog.

**Part C** Add an interface named yelper with a single method called yelp. Declare a value of type yelper and assign the address of the value of type dog. Call the method yelp.

**Part D** Implement the yelper interface for the dog type. Be creative with the bark field. Call the method yelp again from the value of type yelper.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/a-Nzng_E6Z)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/hvVA4zB9Bf))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
