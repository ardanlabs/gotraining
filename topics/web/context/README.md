## Web - Context

Being able to hold onto context while moving from different pieces of middleware down to the final handler can be really important. We'll look at a couple of different ways of handling this data.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  
http://www.gorillatoolkit.org/pkg/context  

## Code Review

Standard Library: [Code](example1/main.go) | [Test](example1/main_test.go)  
Gorilla Context: [Code](example2/main.go) | [Test](example2/main_test.go)  

## Exercises

### Exercise 1

TBD
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
