## Web - Middleware

Learn the basics of using and applying middleware.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  
github.com/urfave/negroni  

## Code Review

Basic middleware: [Code](example1/main.go) | [Test](example1/main_test.go)  
Negroni router: [Code](example2/main.go) | [Test](example2/main_test.go)  

## Exercises

### Exercise 1

Take the Negroni code from example 2 and extend the code by adding a new middleware handler to validate authentication. This call must happen before processing anything else. If authentication fails return a 500 and cancel the processing of the request. If authentication succeeds finish processing the rest of the handlers. Use a query string to cause authentication to succeed or fail.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
