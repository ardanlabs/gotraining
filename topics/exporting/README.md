## Exporting

Packages contain the basic unit of compiled code. They define a scope for the indentifiers that are declared within them. Learning how to package our code is vital because exported identifiers become part of the packages API. Stable and useable API's are incredibily important.

## Notes

* Code in go is complied into packages and then linked together.
* Identifiers are exported (or remain unexported) based on letter-case.
* We import packages to access exported identifiers.
* Any package can use a value of an unexported type.

## Links

http://blog.golang.org/organizing-go-code

http://www.goinggo.net/2014/03/exportedunexported-identifiers-in-go.html

http://www.goinggo.net/2013/08/organizing-code-to-support-go-get.html

##  Package Design

Sandi Metz : Less - The Path to Better Design:  
https://vimeo.com/26330100

Think of a package as a component that accepts a set of inputs and provides a set of predictable outputs. It provides a semantic set of functionality. Focus more on the behavior and less on the concrete.

Package dependency is a choice and must be thought about during the design of the application. Design the source tree and import choices as part of the design of the architecture.

Packages must only depend on other packages that are more stable than itself.

If a package will be a dependency for other packages, it needs to have an API that reveals as little as possible about the package.

Uncertainty is not a license to guess but a directive to decouple.

A package API must consider how changes to itself will affect changes to those packages that depend on it. We must recognize and minimize cascading changes.

When dependencies are weakened and the coupling is loosened, stability is improved and cascading changes are minimized. Allowing for architectures that can better adapt to change over time.

Don’t guess what changes could come, guess what could change.

Knowledge creates dependencies, unstable dependencies increase risk, uncertainty is your guide, loosen coupling so you can reduce cost.

Interfaces provide the highest form of decoupling when the code working with the interfaces does not have any knowledge of the concrete type values stored within them. The concrete types should be opaque to the code using the interfaces. 

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
