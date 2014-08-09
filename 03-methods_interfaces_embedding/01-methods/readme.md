## Methods - Methods, Interfaces and Embedding

Methods are functions that are declared with a receiver which binds the method to a type. Then the method and can be used to operate on values or pointers of that type.

## Notes

* Methods are functions that contain a receiver value.
* Receivers bind a method to a type and can be value or pointers.
* Methods are called against values and pointers, not packages.

## Code Review

[Declare and receiver behavior](example1/example1.go) ([Go Playground](http://play.golang.org/p/jpf5IrVIxf))

[Named typed methods](example2/example2.go) ([Go Playground](http://play.golang.org/p/KKttmFKTVg))

## Exercises

### Exercise 1

Declare a struct that represents a baseball player. Include name, atBats and hits. Declare a method that calculates a players batting average. The formula is Hits / AtBats. Declare a slice of this type and initalize the slice with several players. Iterate over the slice displaying the players name and batting average.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](NEED PLAYGROUND))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)