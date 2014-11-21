## Variables - Language Syntax

Variables are at the heart of the langugage and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type serious and will not allow us to use variables outside the scope of how they are declared.

## Notes

* When variables are being declared to their zero value, use the keyword var.
* When variables are being declared and initialized, use the short variable declaration opertator.
* Escape analysis is used to determine when a value escapes to the heap.

## Links

[Built-In Types](http://golang.org/ref/spec#Boolean_types)

https://golang.org/doc/effective_go.html#variables

http://www.goinggo.net/2013/08/gustavos-ieee-754-brain-teaser.html

## Code Review

[Declare and initalize variables](example1/example1.go) ([Go Playground](http://play.golang.org/p/m4PJ0FpSwX))

## Exercises

### Exercise 1 

**Part A:** Declare three variables that are initalized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

**Part B:** Declare a new variable of type float32 and initalize the variable by converting the literal value of Pi (3.14).

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/Kr7CaO6LdF))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).