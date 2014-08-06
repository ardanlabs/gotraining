## Interfaces
Interfaces provide a way to declare types that define behavior. Then struct and named type can declare methods and implement this behavior. When a struct or named type implements the behavior declare by an interface, it satisfies the interface and can be assigned as values of the interface type. This leads to providing polymorphic behavior in our programs.

### Code Review

[Declare, implement and method call restrictions I](example1/example1.go) ([Go Playground](http://play.golang.org/p/YXhZE1HPUH))

[Declare, implement and method call restrictions II](example2/example2.go) ([Go Playground](http://play.golang.org/p/TEK2rfDrNx))

### Exercise 1
Declare an interface named Speaker with a method named SayHello. Declare a struct named English that represents a person who speaks english and declare a struct named Chinese for someone who speaks chinese. Implement the Speaker interface for each struct using a pointer receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type Speaker and assign the _address of_ a value of type English and call the method. Do it again for a value of type Chinese.

### Exercise 2
From exercise 1, add a new function named SayHello that accepts a value of type Speaker. Implement that function to call the SayHello method on the interface value. Then change the program to pass the address of each struct type to the function.

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)