## Packaging

Packages contain the basic unit of code. All code is built into packages that can be imported into and shared between projects. The Go toolset seeing packaging as a very important part of the dependency management story with tool like "go get", which can pull down source code for projects that live in version control systems and also pull all the dependent code at the same time. Learning how to package our code is vital to building robust and scalable code in Go.

## Notes

* Code in go is complied into packages and then linked together.
* Identifiers are either exported or unexported from a package.
* We import packages to access exported identifiers.
* Any package can use a value of an unexported type.

## Links

http://blog.golang.org/organizing-go-code

http://www.goinggo.net/2014/03/exportedunexported-identifiers-in-go.html

http://www.goinggo.net/2013/08/organizing-code-to-support-go-get.html

## Code Review

[Declare and access exported identifiers](example1/example1.go)

[Declare unexported identifiers and restrictions](example2/example2.go)

[Access values of unexported identifiers](example3/example3.go)

[Unexported struct type fields](example4/example4.go)

[Exported embedded types](example5/example5.go)

[unexported embedded types](example6/example6.go)

## Exercises

### Exercise 1
**Part A** Create a package named toy with a single unexported struct type named bat. Add the exported fields Height and Weight to the bat type. Then create an exported factory method called NewBat that returns pointers of type bat that are initialized to their zero value.

**Part B** Create a program that imports the toy package. Use the NewBat function to create a value of bat and populate the values of Height and Width. Then display the value of the bat variable.

[Answer](exercises/exercise1/exercise1.go)

___
[![GoingGo Training](../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).