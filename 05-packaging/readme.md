## Packaging

### Code Review

[Declare and access exported identifiers](example1/example1.go)

[Declare unexported identifiers and restrictions](example2/example2.go)

[Access values of unexported identifiers](example3/example3.go)

[Unexported struct type fields](example4/example4.go)

[Exported embedded types](example5/example5.go)

[unexported embedded types](example6/example6.go)

### Exercise 1
Create a package that exports functions that can be used to perform simple math operations (add, subtract, multiplication and division). Write a main function that uses this package to perform and display these mathematical operations. 

### Exercise 2
Create a package with an unexported struct type with exported fields. In the package, declare an exported factory function named New to create and return pointers of this unexported type. Write a main function that uses this package to create values of this type, then initialize and display those values.

___
[![GoingGo Training](../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)