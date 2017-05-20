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

## Exercises

### Exercise 1

Write tests that exercise the different endpoints of the "Hello world" language
exercise from [the basics](../basics/README.md). Do this by combining the mux
with the defined routes with a `httptest.NewServer`.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
