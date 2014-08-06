## Embedding

### Code Review

[Embedding types](example1/example1.go) ([Go Playground](http://play.golang.org/p/AQlYR3zQqw))

[Embedded types and interfaces](example2/example2.go) ([Go Playground](http://play.golang.org/p/8vI4KDm2sG))

[Outer and inner type interface implementations](example3/example3.go) ([Go Playground](http://play.golang.org/p/W89veLizhb))

### Exercise 1
Declare a struct type named Animal with two fields associated with all animals. Declare a struct type named Dog with two field associated with a dog. Embed the Animal type into the Dog type. Declare and initalize a value of type Dog. Display the value of the variable.

### Exercise 2
From exercise 1, add a method to the Animal type using a pointer reciever named Yelp which displays the literal string "Not Implemented". Call the method from the value of type Dog.

### Exercise 3
From exercise 2, add an interface named Speaker with a single method called Yelp. Declare a value of type Speaker and assign the address of the value of type Dog. Call the method Yelp.

### Exercise 4
From exercise 3, implement the Speaker interface for the Dog type. Call the method Yelp again from the value of type Speaker.

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)