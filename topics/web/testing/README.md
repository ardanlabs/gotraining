## Web - Testing

Learn the basics of testing web services and applications in Go.

## Notes

* The standard library has a package named httptest with good support.
* There are several ways to create unit and integration tests in Go.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  

## Code Review

[Basic Unit Testing](example1/unit_test.go)  
[Using a http.Handler](example2/unit_test.go)  
[Testing Routes](example3/unit_test.go)  
[Mocking Servers](example4/integration_test.go)  
[Mocking Servers With http.Handler](example5/integration_test.go)  

## Exercises

### Exercise 1

Write a simple web service that has a set of different routes that return the string "Hello World" in multiple languages. Build the service using an Application context that will own the different handler methods. Then create your own mux, bind the routes and start the service. Validate your routes work in your browser.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
