## Exporting

Packages contain the basic unit of compiled code. They define a scope for the indentifiers that are declared within them. Learning how to package our code is vital because exported identifiers become part of the packages API. Stable and useable API's are incredibily important.

## Notes

* Code in go is complied into packages and then linked together.
* Identifiers are exported (or remain unexported) based on letter-case.
* We import packages to access exported identifiers.
* Any package can use a value of an unexported type.

## Links

http://blog.golang.org/organizing-go-code

http://blog.golang.org/package-names

http://www.goinggo.net/2014/03/exportedunexported-identifiers-in-go.html

http://www.goinggo.net/2013/08/organizing-code-to-support-go-get.html

##  Package Design

Sandi Metz : Less - The Path to Better Design:  
https://vimeo.com/26330100

* If you don't understand the data you are working with, you don't understand the problem you are trying to solve.

* Every problem is a data transformation problem at heart and each function, method and workflow must focus on implementing their specific data transformation.

* If your data is changing, your problem is changing.

* When your problem is changing, the data transformations you have implemented need to change.

* Recognizing and minimizing cascading changes across different packages is a way to architect adaptability and stability in your software.

* Packages in Go provide API boundaries that should focus on solving one specific problem or a highly focused group of problems.

* In many languages folders are used to organize code, in Go folders are used to organize API's.

* You must do your best to guess what data could change over time and consider how these changes will affect the software.

* You must understand how changes to the data for a particular package affects the other packages that depend on it.

* When dependencies between packages are weakened and the coupling loosened, cascading changes are minimized and stability is improved.

* Uncertainty about the data is not a license to guess but a directive to decouple.

* Decoupling means reducing the amount of intimate knowledge packages must have about each other to be used together.

* Interfaces provide the highest form of decoupling when the concrete types used to implement them can remain opaque.

## Code Review

[Declare and access exported identifiers](example1/example1.go)

[Declare unexported identifiers and restrictions](example2/example2.go)

[Access values of unexported identifiers](example3/example3.go)

[Unexported struct type fields](example4/example4.go)

[Unexported embedded types](example5/example5.go)

## Exercises

### Exercise 1
**Part A** Create a package named toy with a single exported struct type named Toy. Add the exported fields Name and Weight. Then add two unexported fields named onHand and sold. Declare a factory function called New to create values of type toy and accept parameters for the exported fields. Then declare methods that return and update values for the unexported fields.

**Part B** Create a program that imports the toy package. Use the New function to create a value of type toy. Then use the methods to set the counts and display the field values of that toy value.

[Template](exercises/template1) | 
[Answer](exercises/exercise1)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
