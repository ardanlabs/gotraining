## Web - Sessions and Cookies

Learn the basics of working with session and cookies for web applications.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  
https://github.com/gorilla/sessions  

## Code Review

Using Session: [Code](example1/main.go) | [Test](example1/main_test.go)  
Using Cookies: [Code](example2/main.go) | [Test](example2/main_test.go)  

## Exercises

### Exercise 1

Take the session example and add support for storing the persons phone number. Save the information in session and make sure it can be retrieved. Then update the tests to validate this is working.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
