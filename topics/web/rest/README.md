## Web - REST

Learn the basics of using regex to provide support for building a REST based API in Go.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  

## Code Review

REST With regex: [Code](example1/main.go) | [Test](example1/main_test.go)  

## Exercises

### Exercise 1

Take the CRUD code from example 1 (REST) and extend the code by adding a `PUT` and `DELETE` route. Make sure the routes for both calls ask for the `id` of the customer. Write two new handler functions and bind them into the service so they can be processed. Finally add tests to validate the new routes are working. For both calls redirect the user back to the index page.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).