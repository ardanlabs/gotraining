## Packaging
Packages contain the basic unit of code. All code is built into packages that can be imported into and shared between projects. The Go toolset seeing packaging as a very important part of the dependency management story with tool like "go get", which can pull down source code for projects that live in version control systems and also pull all the dependent code at the same time. Learning how to package our code is vital to building robust and scalable code in Go.

### Code Review

[Declare and access exported identifiers](example1/example1.go)

[Declare unexported identifiers and restrictions](example2/example2.go)

[Access values of unexported identifiers](example3/example3.go)

[Unexported struct type fields](example4/example4.go)

[Exported embedded types](example5/example5.go)

[unexported embedded types](example6/example6.go)

### Exercises

#### Exercise 1
Create a package that exports functions that can be used to perform simple math operations (add, subtract, multiplication and division). Write a main function that uses this package to perform and display these mathematical operations. 

#### Exercise 2
Create a package with an unexported struct type with exported fields. In the package, declare an exported factory function named New to create and return pointers of this unexported type. Write a main function that uses this package to create values of this type, then initialize and display those values.

___
[![GoingGo Training](../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)