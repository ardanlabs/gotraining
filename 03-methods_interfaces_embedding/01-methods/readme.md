## Methods
Methods are functions that are declared with a receiver which binds the method to a type. Then the method and can be used to operate on values of that type. Methods also play a role with the polymorphic behavior we can implement in our programs.

### Code Review

[Declare and receiver behavior](example1/example1.go) ([Go Playground](http://play.golang.org/p/jpf5IrVIxf))

[Named typed methods](example2/example2.go) ([Go Playground](http://play.golang.org/p/KKttmFKTVg))

### Exercise 1
Declare a struct type to maintain information about a hobby. Declare a method that displays a pretty print view of the type's fields. Declare and intialize a value of this type and use the method to display its value.

### Exercise 2
Declare a struct that represents a baseball player. Include Name, AtBats and Hits. Declare a method that calculates a players batting average. The formula is Hits / AtBats. Declare a slice of this type and initalize the slice with several players. Iterate over the slice displaying the players name and batting average.

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)