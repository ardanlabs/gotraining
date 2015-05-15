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

**Part A** Declare a nail type using an empty struct. Then declare two interfaces, one named nailDriver and the other named nailPuller. nailDriver has one method named driveNail that accepts a slice of nails and returns a slice of nails. nailPuller has one method named pullNail that also accepts and returns a slice of nails,
but also accepts the totalNails that exist.

**Part B** Declare a tool type named clawhammer using an empty struct. Implement the two interfaces using a value receiver. For the driveNail method, check that
the number of nails in the slice is not zero and then remove the first element of the slice. For the pullNail method, check the number of nails in the slice with the total nails and then append a new nail to the slice.

**Part C** Declare a toolbox type with two fields named totalNails and nails. The totalNails field will be of type int and the nails field will be a slice of nails. Then embed the clawhammer tool inside the toolbox type. Declare a method named addNails using a pointer reciever that accepts an integer for the number
of nails to add to the toolbox. Append that number of nail values to the slice and set the totalNails field. Declare a method named nailCount using a pointer receiver that returns the total number of nails in the slice and the value for the totalNails field.

**Part D** Write a main function that creates a toolbox, adds nails to the toolbox, displays the nail count, uses the clawhammer to fasten some nails, displays
the nail count again, unfastens the nails and displays the nail count one more time.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/XOBHUvz5uz)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/aIES0zfHfg))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
