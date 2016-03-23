## Packaging / Exporting

Packages contain the basic unit of compiled code. They define a scope for the indentifiers that are declared within them. Learning how to package our code is vital because exported identifiers become part of the packages API. Stable and useable API's are incredibily important.

## Notes

* Code in go is complied into packages and then linked together.
* Identifiers are exported (or remain unexported) based on letter-case.
* We import packages to access exported identifiers.
* Any package can use a value of an unexported type.
* Package design and understanding the components you need is critical.

## Links

http://blog.golang.org/organizing-go-code

http://blog.golang.org/package-names

http://www.goinggo.net/2014/03/exportedunexported-identifiers-in-go.html

http://www.goinggo.net/2013/08/organizing-code-to-support-go-get.html

## An Interview with Brian Kernighan

http://www.cs.cmu.edu/~mihaib/kernighan-interview/index.html

_I think that the real problem with C is that it doesn't give you enough mechanisms for structuring really big programs, for creating ``firewalls'' within programs so you can keep the various pieces apart. It's not that you can't do all of these things, that you can't simulate object-oriented programming or other methodology you want in C. You can simulate it, but the compiler, the language itself isn't giving you any help._ - July 2000

## Design Guidelines

* Learn about the [design guidelines](../../reading/design_guidelines.md) for packaging.

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
