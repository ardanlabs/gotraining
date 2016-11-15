## Web - Templates

Learn the basics of using templates to parse and generate markup.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  
https://github.com/GeertJohan/go.rice  

## Code Review

Basic Template: [Code](example1/main.go) | [Test](example1/main_test.go)  
Data Parsing: [Code](example2/main.go) | [Test](example2/main_test.go)  
Struct Parsing: [Code](example3/main.go) | [Test](example3/main_test.go)  
Escaping: [Code](example4/main.go) | [Test](example4/main_test.go)  
Escaping 2: [Code](example5/main.go) | [Test](example5/main_test.go)  
Complex Markup: [Code](example6/main.go) | [Test](example6/main_test.go)  
Serving Assets: [Code](example7/main.go) | [Test](example7/main_test.go)  
Bundling Assets: [Code](example8/main.go) | [Test](example8/main_test.go) | [Assets](example8/rice-box.go)  

## Exercises

### Exercise 1

Take the code from example 5 (Complex Markup) and add a map of key/value pairs that represents the user's roles. Then add a section to the template to render the set of roles a user has.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
