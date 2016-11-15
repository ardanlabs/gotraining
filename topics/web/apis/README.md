## Web - API

Learn the basics of building a web API in Go.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  

## Code Review

Route Versioning: [Code](example1/main.go) | [Test](example1/main_test.go)  
Header Versioning: [Code](example2/main.go) | [Test](example2/main_test.go)  
CRUD: [Code](example3/main.go) | [Test](example3/main_test.go)  
Echo: [Code](example4/main.go) | [Test](example4/main_test.go)  

## Exercises

### Exercise 1

Take the CRUD code from example 3 (CRUD) and extend the code by adding a `PUT` and `DELETE` route. Make sure the routes for both calls ask for the `id` of the customer. Write two new handler functions and bind them into the service so they can be processed. Finally add tests to validate the new routes are working.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
